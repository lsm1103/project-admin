<script>
    import Field from "@/components/base/Field.svelte";

    export let key = "";
    export let options = {};
</script>

<div class="grid">
    <div class="col-sm-6">
        <Field class="form-field" name="schema.{key}.options.min" let:uniqueId>
            <label for={uniqueId}>最小</label>
            <input type="number" id={uniqueId} bind:value={options.min} />
        </Field>
    </div>

    <div class="col-sm-6">
        <Field class="form-field" name="schema.{key}.options.max" let:uniqueId>
            <label for={uniqueId}>最大</label>
            <input type="number" id={uniqueId} min={options.min} bind:value={options.max} />
        </Field>
    </div>
</div>
