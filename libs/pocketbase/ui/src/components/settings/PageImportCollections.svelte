<script>
    import { tick } from "svelte";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import { pageTitle } from "@/stores/app";
    import { addErrorToast } from "@/stores/toasts";
    import { setErrors } from "@/stores/errors";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import Field from "@/components/base/Field.svelte";
    import SettingsSidebar from "@/components/settings/SettingsSidebar.svelte";
    import ImportPopup from "@/components/settings/ImportPopup.svelte";

    $pageTitle = "导入集合";

    let fileInput;
    let importPopup;

    let schemas = "";
    let isLoadingFile = false;
    let newCollections = [];
    let oldCollections = [];
    let deleteMissing = true;
    let collectionsToUpdate = [];
    let isLoadingOldCollections = false;

    $: if (typeof schemas !== "undefined") {
        loadNewCollections(schemas);
    }

    $: isValid =
        !!schemas &&
        newCollections.length &&
        newCollections.length === newCollections.filter((item) => !!item.id && !!item.name).length;

    $: collectionsToDelete = oldCollections.filter((collection) => {
        return isValid && deleteMissing && !CommonHelper.findByKey(newCollections, "id", collection.id);
    });

    $: collectionsToAdd = newCollections.filter((collection) => {
        return isValid && !CommonHelper.findByKey(oldCollections, "id", collection.id);
    });

    $: if (typeof newCollections !== "undefined" || typeof deleteMissing !== "undefined") {
        loadCollectionsToUpdate();
    }

    $: hasChanges =
        !!schemas && (collectionsToDelete.length || collectionsToAdd.length || collectionsToUpdate.length);

    $: canImport = !isLoadingOldCollections && isValid && hasChanges;

    $: idReplacableCollections = newCollections.filter((collection) => {
        let old =
            CommonHelper.findByKey(oldCollections, "name", collection.name) ||
            CommonHelper.findByKey(oldCollections, "id", collection.id);

        if (!old) {
            return false; // new
        }

        if (old.id != collection.id) {
            return true;
        }

        // check for matching schema fields
        const oldSchema = Array.isArray(old.schema) ? old.schema : [];
        const newSchema = Array.isArray(collection.schema) ? collection.schema : [];
        for (const field of newSchema) {
            const oldFieldById = CommonHelper.findByKey(oldSchema, "id", field.id);
            if (oldFieldById) {
                continue; // no need to do any replacements
            }

            const oldFieldByName = CommonHelper.findByKey(oldSchema, "name", field.name);
            if (oldFieldByName && field.id != oldFieldByName.id) {
                return true;
            }
        }

        return false;
    });

    loadOldCollections();

    async function loadOldCollections() {
        isLoadingOldCollections = true;

        try {
            oldCollections = await ApiClient.collections.getFullList(200);
            // delete timestamps
            for (let collection of oldCollections) {
                delete collection.created;
                delete collection.updated;
            }
        } catch (err) {
            ApiClient.errorResponseHandler(err);
        }

        isLoadingOldCollections = false;
    }

    function loadCollectionsToUpdate() {
        collectionsToUpdate = [];

        if (!isValid) {
            return;
        }

        for (let newCollection of newCollections) {
            const oldCollection = CommonHelper.findByKey(oldCollections, "id", newCollection.id);
            if (
                // no old collection
                !oldCollection?.id ||
                // no changes
                !CommonHelper.hasCollectionChanges(oldCollection, newCollection, deleteMissing)
            ) {
                continue;
            }

            collectionsToUpdate.push({
                new: newCollection,
                old: oldCollection,
            });
        }
    }

    function loadNewCollections() {
        newCollections = [];

        try {
            newCollections = JSON.parse(schemas);
        } catch (_) {}

        if (!Array.isArray(newCollections)) {
            newCollections = [];
        } else {
            newCollections = CommonHelper.filterDuplicatesByKey(newCollections);
        }

        // normalizations
        for (let collection of newCollections) {
            // delete timestamps
            delete collection.created;
            delete collection.updated;

            // merge fields with duplicated ids
            collection.schema = CommonHelper.filterDuplicatesByKey(collection.schema);
        }
    }

    function replaceIds() {
        for (let collection of newCollections) {
            const old =
                CommonHelper.findByKey(oldCollections, "name", collection.name) ||
                CommonHelper.findByKey(oldCollections, "id", collection.id);

            if (!old) {
                continue;
            }

            const originalId = collection.id;
            const replacedId = old.id;
            collection.id = replacedId;

            // replace field ids
            const oldSchema = Array.isArray(old.schema) ? old.schema : [];
            const newSchema = Array.isArray(collection.schema) ? collection.schema : [];
            for (const field of newSchema) {
                const oldField = CommonHelper.findByKey(oldSchema, "name", field.name);
                if (oldField && oldField.id) {
                    field.id = oldField.id;
                }
            }

            // update references
            for (let ref of newCollections) {
                if (!Array.isArray(ref.schema)) {
                    continue;
                }
                for (let field of ref.schema) {
                    if (field.options?.collectionId && field.options?.collectionId === originalId) {
                        field.options.collectionId = replacedId;
                    }
                }
            }
        }

        schemas = JSON.stringify(newCollections, null, 4);
    }

    function loadFile(file) {
        isLoadingFile = true;

        const reader = new FileReader();

        reader.onload = async (event) => {
            isLoadingFile = false;
            fileInput.value = ""; // reset

            schemas = event.target.result;

            await tick();

            if (!newCollections.length) {
                addErrorToast("不正确的集合配置");
                clear();
            }
        };

        reader.onerror = (err) => {
            console.warn(err);
            addErrorToast("获取json失败");

            isLoadingFile = false;
            fileInput.value = ""; // reset
        };

        reader.readAsText(file);
    }

    function clear() {
        schemas = "";
        fileInput.value = "";
        setErrors({});
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
            {#if isLoadingOldCollections}
                <div class="loader" />
            {:else}
                <input
                    bind:this={fileInput}
                    type="file"
                    class="hidden"
                    accept=".json"
                    on:change={() => {
                        if (fileInput.files.length) {
                            loadFile(fileInput.files[0]);
                        }
                    }}
                />

                <div class="content txt-xl m-b-base">
                    <p>
                        粘贴您要导入的集合配置至下方，或
                        <button
                            class="btn btn-outline btn-sm m-l-5"
                            class:btn-loading={isLoadingFile}
                            on:click={() => {
                                fileInput.click();
                            }}
                        >
                            <span class="txt">从json文件中载入</span>
                        </button>
                    </p>
                </div>

                <Field class="form-field {!isValid ? 'field-error' : ''}" name="collections" let:uniqueId>
                    <label for={uniqueId} class="p-b-10">集合</label>
                    <textarea
                        id={uniqueId}
                        class="code"
                        spellcheck="false"
                        rows="15"
                        required
                        bind:value={schemas}
                    />

                    {#if !!schemas && !isValid}
                        <div class="help-block help-block-error">错误的集合配置</div>
                    {/if}
                </Field>

                {#if false}
                    <!-- for now hide the delete control and eventually enable/remove based on the users feedback -->
                    <Field class="form-field form-field-toggle" let:uniqueId>
                        <input
                            type="checkbox"
                            id={uniqueId}
                            bind:checked={deleteMissing}
                            disabled={!isValid}
                        />
                        <label for={uniqueId}>删除失去的集合和字段</label>
                    </Field>
                {/if}

                {#if isValid && newCollections.length && !hasChanges}
                    <div class="alert alert-info">
                        <div class="icon">
                            <i class="ri-information-line" />
                        </div>
                        <div class="content">
                            <string>您的集合配置以及是最新的了</string>
                        </div>
                    </div>
                {/if}

                {#if isValid && newCollections.length && hasChanges}
                    <h5 class="section-title">检测到的更改</h5>

                    <div class="list">
                        {#if collectionsToDelete.length}
                            {#each collectionsToDelete as collection (collection.id)}
                                <div class="list-item">
                                    <span class="label label-danger list-label">已删除</span>
                                    <strong>{collection.name}</strong>
                                    {#if collection.id}
                                        <small class="txt-hint">({collection.id})</small>
                                    {/if}
                                </div>
                            {/each}
                        {/if}

                        {#if collectionsToUpdate.length}
                            {#each collectionsToUpdate as pair (pair.old.id + pair.new.id)}
                                <div class="list-item">
                                    <span class="label label-warning list-label">改变</span>
                                    <div class="inline-flex flex-gap-5">
                                        {#if pair.old.name !== pair.new.name}
                                            <strong class="txt-strikethrough txt-hint">{pair.old.name}</strong
                                            >
                                            <i class="ri-arrow-right-line txt-sm" />
                                        {/if}
                                        <strong>
                                            {pair.new.name}
                                        </strong>
                                        {#if pair.new.id}
                                            <small class="txt-hint">({pair.new.id})</small>
                                        {/if}
                                    </div>
                                </div>
                            {/each}
                        {/if}

                        {#if collectionsToAdd.length}
                            {#each collectionsToAdd as collection (collection.id)}
                                <div class="list-item">
                                    <span class="label label-success list-label">新增</span>
                                    <strong>{collection.name}</strong>
                                    {#if collection.id}
                                        <small class="txt-hint">({collection.id})</small>
                                    {/if}
                                </div>
                            {/each}
                        {/if}
                    </div>
                {/if}

                {#if idReplacableCollections.length}
                    <div class="alert alert-warning m-t-base">
                        <div class="icon">
                            <i class="ri-error-warning-line" />
                        </div>
                        <div class="content">
                            <string>
                                某些导入的集合共享相同的名称和/或字段，但为：使用不同的ID导入。如果需要，可以在导入中替换它们，到
                            </string>
                        </div>
                        <button
                            type="button"
                            class="btn btn-warning btn-sm btn-outline"
                            on:click={() => replaceIds()}
                        >
                            <span class="txt">用原始ids替换</span>
                        </button>
                    </div>
                {/if}

                <div class="flex m-t-base">
                    {#if !!schemas}
                        <button type="button" class="btn btn-secondary link-hint" on:click={() => clear()}>
                            <span class="txt">清除</span>
                        </button>
                    {/if}
                    <div class="flex-fill" />
                    <button
                        type="button"
                        class="btn btn-expanded btn-warning m-l-auto"
                        disabled={!canImport}
                        on:click={() => importPopup?.show(oldCollections, newCollections, deleteMissing)}
                    >
                        <span class="txt">保存</span>
                    </button>
                </div>
            {/if}
        </div>
    </div>
</PageWrapper>

<ImportPopup bind:this={importPopup} on:submit={() => clear()} />

<style>
    .list-label {
        min-width: 65px;
    }
</style>
