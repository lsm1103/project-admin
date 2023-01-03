<script>
    import { Collection } from "pocketbase";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import CodeBlock from "@/components/base/CodeBlock.svelte";
    import SdkTabs from "@/components/collections/docs/SdkTabs.svelte";

    export let collection = new Collection();

    let responseTab = 200;
    let responses = [];

    $: adminsOnly = collection?.createRule === null;

    $: backendAbsUrl = CommonHelper.getApiExampleUrl(ApiClient.baseUrl);

    $: responses = [
        {
            code: 200,
            body: JSON.stringify(CommonHelper.dummyCollectionRecord(collection), null, 2),
        },
        {
            code: 400,
            body: `
                {
                  "code": 400,
                  "message": "Failed to create record.",
                  "data": {
                    "${collection?.schema?.[0]?.name}": {
                      "code": "validation_required",
                      "message": "Missing required value."
                    }
                  }
                }
            `,
        },
        {
            code: 403,
            body: `
                {
                  "code": 403,
                  "message": "You are not allowed to perform this request.",
                  "data": {}
                }
            `,
        },
    ];
</script>

<div class="alert alert-success">
    <strong class="label label-primary">POST</strong>
    <div class="content">
        <p>
            /api/collections/<strong>{collection.name}</strong>/records
        </p>
    </div>
    {#if adminsOnly}
        <p class="txt-hint txt-sm txt-right">需要 <code>Authorization: Admin TOKEN</code> header</p>
    {/if}
</div>

<div class="content m-b-base">
    <p>新建一个新的属于<strong>{collection.name}</strong>的记录</p>
    <p>
        正文参数可以作为 <code>application/json</code> 或者
        <code>multipart/form-data</code>发送
    </p>
    <p>
        文件上传仅支持<code>multipart/form-data</code>.
    </p>
</div>

<div class="section-title">Client SDKs example</div>
<SdkTabs
    js={`
        import PocketBase from 'pocketbase';

        const client = new PocketBase('${backendAbsUrl}');

        ...

        const data = { ... };

        const record = await client.records.create('${collection?.name}', data);
    `}
    dart={`
        import 'package:pocketbase/pocketbase.dart';

        final client = PocketBase('${backendAbsUrl}');

        ...

        final body = <String, dynamic>{ ... };

        final record = await client.records.create('${collection?.name}', body: body);
    `}
/>

<div class="section-title">正文参数</div>
<table class="table-compact table-border m-b-lg">
    <thead>
        <tr>
            <th>参数名</th>
            <th>类型</th>
            <th width="50%">描述</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>
                <div class="inline-flex">
                    <span class="label label-warning">可选的</span>
                    <span>id</span>
                </div>
            </td>
            <td>
                <span class="label">String</span>
            </td>
            <td>
                存储为记录 ID 的<strong>15 个字符串</strong>。
                <br />
                如果未设置，它将自动生成。
            </td>
        </tr>
        {#each collection?.schema as field (field.name)}
            <tr>
                <td>
                    <div class="inline-flex">
                        {#if field.required}
                            <span class="label label-success">必须</span>
                        {:else}
                            <span class="label label-warning">可选</span>
                        {/if}
                        <span>{field.name}</span>
                    </div>
                </td>
                <td>
                    <span class="label">{CommonHelper.getFieldValueType(field)}</span>
                </td>
                <td>
                    {#if field.type === "text"}
                        纯文本值
                    {:else if field.type === "number"}
                        数字值
                    {:else if field.type === "json"}
                        json数组或对象
                    {:else if field.type === "email"}
                        Email地址
                    {:else if field.type === "url"}
                        URL地址
                    {:else if field.type === "file"}
                        文件对象<br />
                        设置为<code>null</code>可以删除已经上传的文件
                    {:else if field.type === "relation"}
                        关系记录 {field.options?.maxSelect > 1 ? "ids" : "id"}.
                    {:else if field.type === "user"}
                        用户 {field.options?.maxSelect > 1 ? "ids" : "id"}.
                    {/if}
                </td>
            </tr>
        {/each}
    </tbody>
</table>

<div class="section-title">路径参数</div>
<table class="table-compact table-border m-b-lg">
    <thead>
        <tr>
            <th>参数名</th>
            <th>类型</th>
            <th width="60%">描述</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>expand</td>
            <td>
                <span class="label">String</span>
            </td>
            <td>
                返回创建的记录时自动展开关系。例如：
                <CodeBlock
                    content={`
                        ?expand=rel1,rel2.subrel21.subrel22
                    `}
                />
                最多支持 6 级深度嵌套关系扩展。 <br />
                扩展的关系将附加到
                <code>@expand</code> 属性下的记录 (eg. <code>{`"@expand": {"rel1": {...}, ...}`}</code>). 
                只有用户有权 <strong>查看</strong> 的关系才会被展开。
            </td>
        </tr>
    </tbody>
</table>

<div class="section-title">响应</div>
<div class="tabs">
    <div class="tabs-header compact left">
        {#each responses as response (response.code)}
            <button
                class="tab-item"
                class:active={responseTab === response.code}
                on:click={() => (responseTab = response.code)}
            >
                {response.code}
            </button>
        {/each}
    </div>
    <div class="tabs-content">
        {#each responses as response (response.code)}
            <div class="tab-item" class:active={responseTab === response.code}>
                <CodeBlock content={response.body} />
            </div>
        {/each}
    </div>
</div>
