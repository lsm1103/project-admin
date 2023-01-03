<script>
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import Field from "@/components/base/Field.svelte";
    import ObjectSelect from "@/components/base/ObjectSelect.svelte";

    export let key = "";
    export let options = {};

    const defaultOptions = [
        { label: "False", value: false },
        { label: "True", value: true },
    ];

    let isLoading = false;
    let collections = [];

    // load defaults
    $: if (CommonHelper.isEmpty(options)) {
        options = {
            maxSelect: 1,
            collectionId: null,
            cascadeDelete: false,
        };
    }

    loadCollections();

    function loadCollections() {
        isLoading = true;

        ApiClient.collections.getFullList(200, { sort: "-created" })
            .then((items) => {
                collections = items;
            })
            .catch((err) => {
                ApiClient.errorResponseHandler(err);
            })
            .finally(() => {
                isLoading = false;
            });
    }
</script>

<div class="grid">
    <div class="col-sm-9">
        <Field class="form-field required" name="schema.{key}.options.collectionId" let:uniqueId>
            <label for={uniqueId}>集合</label>
            <ObjectSelect
                searchable={collections.length > 5}
                selectPlaceholder={isLoading ? "Loading..." : "选择集合"}
                noOptionsText="未发现集合"
                selectionKey="id"
                items={collections}
                bind:keyOfSelected={options.collectionId}
            />
        </Field>
    </div>
    <div class="col-sm-3">
        <Field class="form-field required" name="schema.{key}.options.maxSelect" let:uniqueId>
            <label for={uniqueId}>最大选择数</label>
            <input type="number" id={uniqueId} step="1" min="1" required bind:value={options.maxSelect} />
        </Field>
    </div>
    <div class="col-sm-12">
        <Field class="form-field" name="schema.{key}.options.cascadeDelete" let:uniqueId>
            <label for={uniqueId}>删除记录在relation删除时</label>
            <ObjectSelect id={uniqueId} items={defaultOptions} bind:keyOfSelected={options.cascadeDelete} />
        </Field>
    </div>
</div>
