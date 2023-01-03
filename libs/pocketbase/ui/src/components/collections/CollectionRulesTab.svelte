<script>
    import { onMount, tick } from "svelte";
    import { slide } from "svelte/transition";
    import { Collection } from "pocketbase";
    import tooltip from "@/actions/tooltip";
    import Field from "@/components/base/Field.svelte";

    export let collection = new Collection();

    let tempValues = {};
    let showFiltersInfo = false;
    let editorRefs = {};
    let ruleInputComponent;
    let isRuleComponentLoading = false;

    // all supported collection rules in "collection_rule_prop: label" format
    const ruleProps = {
        listRule: "List Action",
        viewRule: "View Action",
        createRule: "Create Action",
        updateRule: "Update Action",
        deleteRule: "Delete Action",
    };

    function isAdminOnly(propVal) {
        return propVal === null;
    }

    async function loadEditorComponent() {
        isRuleComponentLoading = true;
        try {
            ruleInputComponent = (await import("@/components/base/FilterAutocompleteInput.svelte")).default;
        } catch (err) {
            console.warn(err);
            ruleInputComponent = null;
        }
        isRuleComponentLoading = false;
    }

    onMount(() => {
        loadEditorComponent();
    });
</script>

<div class="block m-b-base">
    <div class="flex">
        <p>
            所有规则都服从
            <a href={import.meta.env.PB_RULES_SYNTAX_DOCS} target="_blank" rel="noopener">
                PocketBase筛选格式和操作符
            </a>.
        </p>
        <span
            class="expand-handle txt-sm txt-bold txt-nowrap link-hint"
            on:click={() => (showFiltersInfo = !showFiltersInfo)}
        >
            {showFiltersInfo ? "隐藏可用字段" : "显示可用字段"}
        </span>
    </div>

    {#if showFiltersInfo}
        <div transition:slide|local={{ duration: 150 }}>
            <div class="alert alert-warning m-0">
                <div class="content">
                    <p class="m-b-0">以下记录字段可用:</p>
                    <div class="inline-flex flex-gap-5">
                        <code>id</code>
                        <code>created</code>
                        <code>updated</code>
                        {#each collection.schema as field}
                            {#if field.type === "relation" || field.type === "user"}
                                <code>{field.name}.*</code>
                            {:else}
                                <code>{field.name}</code>
                            {/if}
                        {/each}
                    </div>

                    <hr class="m-t-10 m-b-5" />

                    <p class="m-b-0">
                        可以使用特殊的 <em>@request</em> 过滤器访问请求字段：
                    </p>
                    <div class="inline-flex flex-gap-5">
                        <code>@request.method</code>
                        <code>@request.query.*</code>
                        <code>@request.data.*</code>
                        <code>@request.user.*</code>
                    </div>

                    <hr class="m-t-10 m-b-5" />

                    <p class="m-b-0">
                        您还可以使用 <em
                            >@collection</em
                        > 过滤器添加约束和查询其他集合：
                    </p>
                    <div class="inline-flex flex-gap-5">
                        <code>@collection.ANY_COLLECTION_NAME.*</code>
                    </div>

                    <hr class="m-t-10 m-b-5" />

                    <p>
                        示例规则：
                        <br />
                        <code>@request.user.id!="" && created>"2022-01-01 00:00:00"</code>
                    </p>
                </div>
            </div>
        </div>
    {/if}
</div>

{#if isRuleComponentLoading}
    <div class="txt-center">
        <span class="loader" />
    </div>
{:else}
    {#each Object.entries(ruleProps) as [prop, label] (prop)}
        <hr class="m-t-sm m-b-sm" />
        <div class="rule-block">
            {#if isAdminOnly(collection[prop])}
                <button
                    type="button"
                    class="rule-toggle-btn btn btn-circle btn-outline btn-success"
                    use:tooltip={"解锁并设置自定义规则"}
                    on:click={async () => {
                        collection[prop] = tempValues[prop] || "";
                        await tick();
                        editorRefs[prop]?.focus();
                    }}
                >
                    <i class="ri-lock-unlock-line" />
                </button>
            {:else}
                <button
                    type="button"
                    class="rule-toggle-btn btn btn-circle btn-outline"
                    use:tooltip={"锁定并设置只有admin可访问"}
                    on:click={() => {
                        tempValues[prop] = collection[prop];
                        collection[prop] = null;
                    }}
                >
                    <i class="ri-lock-line" />
                </button>
            {/if}

            <Field
                class="form-field rule-field m-0 {isAdminOnly(collection[prop]) ? 'disabled' : ''}"
                name={prop}
                let:uniqueId
            >
                <label for={uniqueId}>
                    {label} - {isAdminOnly(collection[prop]) ? "仅admin" : "自定义规则"}
                </label>

                <svelte:component
                    this={ruleInputComponent}
                    id={uniqueId}
                    bind:this={editorRefs[prop]}
                    bind:value={collection[prop]}
                    baseCollection={collection}
                    disabled={isAdminOnly(collection[prop])}
                />

                <div class="help-block">
                    {#if isAdminOnly(collection[prop])}
                        只有管​​理员才能访问（解锁更改）
                    {:else}
                        留空以授予所有人访问权限
                    {/if}
                </div>
            </Field>
        </div>
    {/each}
{/if}

<style>
    .rule-block {
        display: flex;
        align-items: flex-start;
        gap: var(--xsSpacing);
    }
    .rule-toggle-btn {
        margin-top: 15px;
    }
</style>
