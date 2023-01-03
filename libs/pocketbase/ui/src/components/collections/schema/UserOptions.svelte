<script>
    import CommonHelper from "@/utils/CommonHelper";
    import Field from "@/components/base/Field.svelte";
    import ObjectSelect from "@/components/base/ObjectSelect.svelte";

    const defaultOptions = [
        { label: "False", value: false },
        { label: "True", value: true },
    ];

    export let key = "";
    export let options = {};

    // load defaults
    $: if (CommonHelper.isEmpty(options)) {
        options = {
            maxSelect: 1,
            cascadeDelete: false,
        };
    }
</script>

<div class="grid">
    <div class="col-sm-6">
        <Field class="form-field required" name="schema.{key}.options.maxSelect" let:uniqueId>
            <label for={uniqueId}>最大选择数</label>
            <input type="number" id={uniqueId} step="1" min="1" required bind:value={options.maxSelect} />
        </Field>
    </div>
    <div class="col-sm-6">
        <Field class="form-field" name="schema.{key}.options.cascadeDelete" let:uniqueId>
            <label for={uniqueId}>在用户删除时删除记录</label>
            <ObjectSelect id={uniqueId} items={defaultOptions} bind:keyOfSelected={options.cascadeDelete} />
        </Field>
    </div>
</div>
