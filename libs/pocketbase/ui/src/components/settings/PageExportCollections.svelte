<script>
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import { pageTitle } from "@/stores/app";
    import { addInfoToast } from "@/stores/toasts";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import CodeBlock from "@/components/base/CodeBlock.svelte";
    import SettingsSidebar from "@/components/settings/SettingsSidebar.svelte";

    $pageTitle = "导出集合";

    const uniqueId = "export_" + CommonHelper.randomString(5);

    let previewContainer;
    let collections = [];
    let isLoadingCollections = false;

    $: schema = JSON.stringify(collections, null, 4);

    loadCollections();

    async function loadCollections() {
        isLoadingCollections = true;

        try {
            collections = await ApiClient.collections.getFullList(100, {
                $cancelKey: uniqueId,
            });
            // delete timestamps
            for (let collection of collections) {
                delete collection.created;
                delete collection.updated;
            }
        } catch (err) {
            ApiClient.errorResponseHandler(err);
        }

        isLoadingCollections = false;
    }

    function download() {
        CommonHelper.downloadJson(collections, "pb_schema");
    }

    function copy() {
        CommonHelper.copyToClipboard(schema);
        addInfoToast("配置已经复制到了剪贴板上", 3000);
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
        <div class="panel">
            {#if isLoadingCollections}
                <div class="loader" />
            {:else}
                <div class="content txt-xl m-b-base">
                    <p>
                        您将在下面找到可以导入另一个 PocketBase 环境的当前集合配置。
                    </p>
                </div>

                <div
                    tabindex="0"
                    bind:this={previewContainer}
                    class="export-preview"
                    on:keydown={(e) => {
                        // select all
                        if (e.ctrlKey && e.code === "KeyA") {
                            e.preventDefault();
                            const selection = window.getSelection();
                            const range = document.createRange();
                            range.selectNodeContents(previewContainer);
                            selection.removeAllRanges();
                            selection.addRange(range);
                        }
                    }}
                >
                    <button
                        type="button"
                        class="btn btn-sm btn-secondary fade copy-schema"
                        on:click={() => copy()}
                    >
                        <span class="txt">复制</span>
                    </button>

                    <CodeBlock content={schema} />
                </div>

                <div class="flex m-t-base">
                    <div class="flex-fill" />
                    <button type="button" class="btn btn-expanded" on:click={() => download()}>
                        <i class="ri-download-line" />
                        <span class="txt">下载json文件</span>
                    </button>
                </div>
            {/if}
        </div>
    </div>
</PageWrapper>

<style>
    .export-preview {
        position: relative;
        height: 500px;
    }
    .export-preview .copy-schema {
        position: absolute;
        right: 15px;
        top: 15px;
    }
</style>
