<script>
    import { slide } from "svelte/transition";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import { pageTitle } from "@/stores/app";
    import { setErrors } from "@/stores/errors";
    import { removeAllToasts, addWarningToast, addSuccessToast } from "@/stores/toasts";
    import tooltip from "@/actions/tooltip";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import Field from "@/components/base/Field.svelte";
    import RedactedPasswordInput from "@/components/base/RedactedPasswordInput.svelte";
    import SettingsSidebar from "@/components/settings/SettingsSidebar.svelte";

    $pageTitle = "文件存储";

    const testRequestKey = "s3_test_request";

    let originalFormSettings = {};
    let formSettings = {};
    let isLoading = false;
    let isSaving = false;
    let isTesting = false;
    let testS3Error = null;
    let testS3TimeoutId = null;

    $: initialHash = JSON.stringify(originalFormSettings);

    $: hasChanges = initialHash != JSON.stringify(formSettings);

    loadSettings();

    async function loadSettings() {
        isLoading = true;

        try {
            const settings = (await ApiClient.settings.getAll()) || {};
            init(settings);
        } catch (err) {
            ApiClient.errorResponseHandler(err);
        }

        isLoading = false;
    }

    async function testS3() {
        testS3Error = null;

        if (!formSettings.s3.enabled) {
            return; // nothing to test
        }

        isTesting = true;

        try {
            await ApiClient.settings.testS3({ $cancelKey: testRequestKey });
        } catch (err) {
            testS3Error = err;
        }

        isTesting = false;
    }

    async function save() {
        if (isSaving || !hasChanges) {
            return;
        }

        isSaving = true;

        // auto cancel the test request after 30sec
        clearTimeout(testS3TimeoutId);
        testS3TimeoutId = setTimeout(() => {
            ApiClient.cancelRequest(testRequestKey);
            addErrorToast("S3测试连接超时");
        }, 30000);

        try {
            ApiClient.cancelRequest(testRequestKey);
            const settings = await ApiClient.settings.update(CommonHelper.filterRedactedProps(formSettings));
            setErrors({});

            await init(settings);

            removeAllToasts();

            if (testS3Error) {
                addWarningToast("成功保存但是未成功建立S3连接");
            } else {
                addSuccessToast("成功保存文件存储设置");
            }
        } catch (err) {
            ApiClient.errorResponseHandler(err);
        }

        clearTimeout(testS3TimeoutId);

        isSaving = false;
    }

    async function init(settings = {}) {
        formSettings = {
            s3: settings?.s3 || {},
        };
        originalFormSettings = JSON.parse(JSON.stringify(formSettings));

        await testS3();
    }

    async function reset() {
        formSettings = JSON.parse(JSON.stringify(originalFormSettings || {}));

        await testS3();
    }
</script>

<SettingsSidebar />

<PageWrapper>
    <header class="page-header">
        <nav class="breadcrumbs">
            <div class="breadcrumb-item">Settings</div>
            <div class="breadcrumb-item">{$pageTitle}</div>
        </nav>
    </header>

    <div class="wrapper">
        <form class="panel" autocomplete="off" on:submit|preventDefault={() => save()}>
            <div class="content txt-xl m-b-base">
                <p>默认情况下，PocketBase 使用本地文件系统来存储上传的文件。</p>
                <p>
                    如果您的磁盘空间有限，您可以选择连接到 S3 兼容存储。
                </p>
            </div>

            {#if isLoading}
                <div class="loader" />
            {:else}
                <Field class="form-field form-field-toggle" let:uniqueId>
                    <input type="checkbox" id={uniqueId} required bind:checked={formSettings.s3.enabled} />
                    <label for={uniqueId}>使用 S3 storage</label>
                </Field>

                {#if originalFormSettings.s3?.enabled != formSettings.s3.enabled}
                    <div transition:slide|local={{ duration: 150 }}>
                        <div class="alert alert-warning m-0">
                            <div class="icon">
                                <i class="ri-error-warning-line" />
                            </div>
                            <div class="content">
                                如果您有现有的上传文件，则必须手动将它们从
                                <strong>
                                    {originalFormSettings.s3?.enabled ? "S3存储" : "本地文件系统"}
                                </strong>
                                迁移到
                                <strong>{formSettings.s3.enabled ? "S3存储" : "本地文件系统"}</strong
                                >。
                                <br />
                                有许多命令行工具可以帮助您，例如：
                                <a
                                    href="https://github.com/rclone/rclone"
                                    target="_blank"
                                    rel="noopener"
                                    class="txt-bold"
                                >
                                    rclone
                                </a>,
                                <a
                                    href="https://github.com/peak/s5cmd"
                                    target="_blank"
                                    rel="noopener"
                                    class="txt-bold"
                                >
                                    s5cmd
                                </a>等
                            </div>
                        </div>
                        <div class="clearfix m-t-base" />
                    </div>
                {/if}

                {#if formSettings.s3.enabled}
                    <div class="grid" transition:slide|local={{ duration: 150 }}>
                        <div class="col-lg-6">
                            <Field class="form-field required" name="s3.endpoint" let:uniqueId>
                                <label for={uniqueId}>Endpoint</label>
                                <input
                                    type="text"
                                    id={uniqueId}
                                    required
                                    bind:value={formSettings.s3.endpoint}
                                />
                            </Field>
                        </div>
                        <div class="col-lg-3">
                            <Field class="form-field required" name="s3.bucket" let:uniqueId>
                                <label for={uniqueId}>Bucket</label>
                                <input
                                    type="text"
                                    id={uniqueId}
                                    required
                                    bind:value={formSettings.s3.bucket}
                                />
                            </Field>
                        </div>
                        <div class="col-lg-3">
                            <Field class="form-field required" name="s3.region" let:uniqueId>
                                <label for={uniqueId}>Region</label>
                                <input
                                    type="text"
                                    id={uniqueId}
                                    required
                                    bind:value={formSettings.s3.region}
                                />
                            </Field>
                        </div>
                        <div class="col-lg-6">
                            <Field class="form-field required" name="s3.accessKey" let:uniqueId>
                                <label for={uniqueId}>Access key</label>
                                <input
                                    type="text"
                                    id={uniqueId}
                                    required
                                    bind:value={formSettings.s3.accessKey}
                                />
                            </Field>
                        </div>
                        <div class="col-lg-6">
                            <Field class="form-field required" name="s3.secret" let:uniqueId>
                                <label for={uniqueId}>Secret</label>
                                <RedactedPasswordInput
                                    id={uniqueId}
                                    required
                                    bind:value={formSettings.s3.secret}
                                />
                            </Field>
                        </div>
                        <div class="col-lg-12">
                            <Field class="form-field" name="s3.forcePathStyle" let:uniqueId>
                                <input
                                    type="checkbox"
                                    id={uniqueId}
                                    bind:checked={formSettings.s3.forcePathStyle}
                                />
                                <label for={uniqueId}>
                                    <span class="txt">强制路径样式寻址</span>
                                    <i
                                        class="ri-information-line link-hint"
                                        use:tooltip={{
                                            text: ' 强制请求使用路径样式寻址 , 例如： 用"https://s3.amazonaws.com/BUCKET/KEY" 来替代 "https://BUCKET.s3.amazonaws.com/KEY".',
                                            position: "top",
                                        }}
                                    />
                                </label>
                            </Field>
                        </div>
                        <!-- margin helper -->
                        <div class="col-lg-12" />
                    </div>
                {/if}

                <div class="flex">
                    <div class="flex-fill" />

                    {#if formSettings.s3?.enabled && !hasChanges && !isSaving}
                        {#if isTesting}
                            <span class="loader loader-sm" />
                        {:else if testS3Error}
                            <div
                                class="label label-sm label-warning entrance-right"
                                use:tooltip={testS3Error.data?.message}
                            >
                                <i class="ri-error-warning-line txt-warning" />
                                <span class="txt">建立S3连接失败</span>
                            </div>
                        {:else}
                            <div class="label label-sm label-success entrance-right">
                                <i class="ri-checkbox-circle-line txt-success" />
                                <span class="txt">S3连接成功建立</span>
                            </div>
                        {/if}
                    {/if}

                    {#if hasChanges}
                        <button
                            type="button"
                            class="btn btn-secondary btn-hint"
                            disabled={isSaving}
                            on:click={() => reset()}
                        >
                            <span class="txt">取消</span>
                        </button>
                    {/if}

                    <button
                        type="submit"
                        class="btn btn-expanded"
                        class:btn-loading={isSaving}
                        disabled={!hasChanges || isSaving}
                        on:click={() => save()}
                    >
                        <span class="txt">保存更改</span>
                    </button>
                </div>
            {/if}
        </form>
    </div>
</PageWrapper>
