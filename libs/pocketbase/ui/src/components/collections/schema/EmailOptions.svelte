<script>
    import CommonHelper from "@/utils/CommonHelper";
    import tooltip from "@/actions/tooltip";
    import Field from "@/components/base/Field.svelte";
    import MultipleValueInput from "@/components/base/MultipleValueInput.svelte";

    export let key = "";
    export let options = {};
</script>

<div class="grid">
    <div class="col-sm-6">
        <Field class="form-field" name="schema.{key}.options.exceptDomains" let:uniqueId>
            <label for={uniqueId}>
                <span class="txt">禁止域名</span>
                <i
                    class="ri-information-line link-hint"
                    use:tooltip={{
                        text: '下列域名都是被禁止的\n 这个字段会被禁用如果设置了允许域名',
                        position: "top",
                    }}
                />
            </label>
            <MultipleValueInput
                id={uniqueId}
                disabled={!CommonHelper.isEmpty(options.onlyDomains)}
                bind:value={options.exceptDomains}
            />
            <div class="help-block">使用,作为分隔符</div>
        </Field>
    </div>

    <div class="col-sm-6">
        <Field class="form-field" name="schema.{key}.options.onlyDomains" let:uniqueId>
            <label for="{uniqueId}.options.onlyDomains">
                <span class="txt">允许域名</span>
                <i
                    class="ri-information-line link-hint"
                    use:tooltip={{
                        text: '下列域名都是被允许的 \n 这个字段会被禁用如果设置了禁止域名',
                        position: "top",
                    }}
                />
            </label>
            <MultipleValueInput
                id="{uniqueId}.options.onlyDomains"
                disabled={!CommonHelper.isEmpty(options.exceptDomains)}
                bind:value={options.onlyDomains}
            />
            <div class="help-block">使用,作为分隔符</div>
        </Field>
    </div>
</div>
