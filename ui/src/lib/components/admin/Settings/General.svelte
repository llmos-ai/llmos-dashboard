<script lang="ts">
  import { onMount, getContext } from "svelte";
  import { updateSettingValue } from "$lib/apis/settings";
  import {toast} from "svelte-sonner";

  const i18n = getContext("i18n");

  export let saveHandler: Function;
  let settings = localStorage.serverSettings ? JSON.parse(localStorage.serverSettings) : [];

  const toggleSettingUpdate = async (name, value) => {
    console.log("toggle update setting:", name, value)
    let res = await updateSettingValue(localStorage.token, name, String(value));
    if (res.error) {
      toast.error(res.error);
      return;
    }
    localStorage.serverSettings = JSON.stringify(settings)
    toast.success(`setting ${name} updated`);
    settings = res
  };

  const onSubmit = function (e) {
    const formData = new FormData(e.target);
    for (let field of formData) {
      const [key, value] = field;
      toggleSettingUpdate(key, value)
      console.log(key)
    }
  }
</script>

<form
  class="flex flex-col h-full justify-between space-y-3 text-sm"
  on:submit|preventDefault={onSubmit}>

  <div class=" space-y-3 pr-1.5 overflow-y-scroll max-h-80">
    <div>
      <div class=" mb-2 text-sm font-medium">{$i18n.t("General Settings")}</div>

      {#each settings as setting}
      {#if setting.name === "signup-enabled"}
      <div class="  flex w-full justify-between">
        <div class=" self-center text-xs font-medium">
          {$i18n.t("Allow New Sign Ups")}
        </div>
        <div class="flex items-center relative">
          <select class="dark:bg-gray-900 w-fit pr-8 rounded py-2 px-2 text-xs bg-transparent outline-none text-right"
                  bind:value={setting.value}
                  on:change={(e) => {
              toggleSettingUpdate(setting.name, e.target.value);
            }}
          >
            <option value="true">{$i18n.t("Allow")}</option>
            <option value="false">{$i18n.t("Don't Allow")}</option>
          </select>
        </div>
      </div>
      {/if}

      {#if setting.name === "default-user-role"}
      <div class=" flex w-full justify-between">
        <div class=" self-center text-xs font-medium">
          {$i18n.t("Default User Role")}
        </div>
        <div class="flex items-center relative">
          <select
            class="dark:bg-gray-900 w-fit pr-8 rounded py-2 px-2 text-xs bg-transparent outline-none text-right"
            bind:value={setting.value}
            placeholder="Select a theme"
            on:change={(e) => {
              toggleSettingUpdate(setting.name, e.target.value);
            }}
          >
            <option value="pending">{$i18n.t("pending")}</option>
            <option value="user">{$i18n.t("user")}</option>
            <option value="admin">{$i18n.t("admin")}</option>
          </select>
        </div>
      </div>
      {/if}

      {#if setting.name === "webhook-url"}
      <hr class=" dark:border-gray-700 my-3" />
      <div class=" w-full justify-between">
        <div class="flex w-full justify-between">
          <div class=" self-center text-xs font-medium">
            {$i18n.t("Webhook URL")}
          </div>
        </div>

        <div class="flex mt-2 space-x-2">
          <input
            class="w-full rounded py-1.5 px-4 text-sm dark:text-gray-300 dark:bg-gray-800 outline-none border border-gray-100 dark:border-gray-600"
            type="text"
            name={setting.name}
            placeholder={`https://example.com/webhook`}
            bind:value={setting.value}
          />
        </div>
      </div>
      {/if}


      {#if setting.name === "token-expire-time"}
      <hr class=" dark:border-gray-700 my-3" />
      <div class=" w-full justify-between">
        <div class="flex w-full justify-between">
          <div class=" self-center text-xs font-medium">
            {$i18n.t("JWT Expiration")}
          </div>
        </div>

        <div class="flex mt-2 space-x-2">
          {#if setting.value}
          <input
            class="w-full rounded py-1.5 px-4 text-sm dark:text-gray-300 dark:bg-gray-800 outline-none border border-gray-100 dark:border-gray-600"
            type="text"
            name={setting.name}
            placeholder={`e.g.) "60m","7h". `}
            bind:value={setting.value}
          />
          {:else}
          <input
            class="w-full rounded py-1.5 px-4 text-sm dark:text-gray-300 dark:bg-gray-800 outline-none border border-gray-100 dark:border-gray-600"
            type="text"
            name={setting.name}
            placeholder={`e.g.) "60m","7h". `}
            bind:value={setting.default}
          />
          {/if}
        </div>

        <div class="mt-2 text-xs text-gray-400 dark:text-gray-500">
          {$i18n.t("Valid time units:")}
          <span class=" text-gray-300 font-medium"
            >{$i18n.t(
              "'s', 'm', 'h'."
            )}</span
          >
        </div>
      </div>
      {/if}
      {/each}
    </div>
  </div>

  <div class="flex justify-end pt-3 text-sm font-medium">
    <button
      class=" px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-gray-100 transition rounded"
      type="submit"
    >
      {$i18n.t("Save")}
    </button>
  </div>
</form>
