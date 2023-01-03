<script>
    import { tick, createEventDispatcher } from "svelte";
    import OverlayPanel from "@/components/base/OverlayPanel.svelte";

    const dispatch = createEventDispatcher();

    let panel;
    let collection;

    $: isCollectionRenamed = collection?.originalName != collection?.name;

    $: renamedFields =
        collection?.schema.filter(
            (field) => field.id && !field.toDelete && field.originalName != field.name
        ) || [];

    $: deletedFields = collection?.schema.filter((field) => field.id && field.toDelete) || [];

    export async function show(collectionToCheck) {
        collection = collectionToCheck;

        await tick();

        if (!isCollectionRenamed && !renamedFields.length && !deletedFields.length) {
            // no confirm required changes
            confirm();
        } else {
            panel?.show();
        }
    }

    export function hide() {
        panel?.hide();
    }

    function confirm() {
        hide();
        dispatch("confirm");
    }
</script>

<OverlayPanel bind:this={panel} class="confirm-changes-panel" popup on:hide on:show>
    <svelte:fragment slot="header">
        <h4>确认集合更改</h4>
    </svelte:fragment>

    <div class="alert alert-warning">
        <div class="icon">
            <i class="ri-error-warning-line" />
        </div>
        <div class="content txt-bold">
            <p>
                如果以下任何更改是另一个集合规则或过滤器的一部分，则您必须手动更新它！
            </p>
            {#if deletedFields.length}
                <p>与删除字段相关的所有数据将被永久删除！</p>
            {/if}
        </div>
    </div>

    <h6>更改:</h6>
    <ul class="changes-list">
        {#if isCollectionRenamed}
            <li>
                <div class="inline-flex">
                    重命名集合
                    <strong class="txt-strikethrough txt-hint">{collection.originalName}</strong>
                    <i class="ri-arrow-right-line txt-sm" />
                    <strong class="txt"> {collection.name}</strong>
                </div>
            </li>
        {/if}

        {#each renamedFields as field}
            <li>
                <div class="inline-flex">
                    重命名字段
                    <strong class="txt-strikethrough txt-hint">{field.originalName}</strong>
                    <i class="ri-arrow-right-line txt-sm" />
                    <strong class="txt"> {field.name}</strong>
                </div>
            </li>
        {/each}

        {#each deletedFields as field}
            <li class="txt-danger">
                删除字段 <span class="txt-bold">{field.name}</span>
            </li>
        {/each}
    </ul>

    <svelte:fragment slot="footer">
        <!-- svelte-ignore a11y-autofocus -->
        <button autofocus type="button" class="btn btn-secondary" on:click={() => hide()}>
            <span class="txt">取消</span>
        </button>
        <button type="button" class="btn btn-expanded" on:click={() => confirm()}>
            <span class="txt">确认</span>
        </button>
    </svelte:fragment>
</OverlayPanel>

<style>
    .changes-list {
        word-break: break-all;
    }
</style>
