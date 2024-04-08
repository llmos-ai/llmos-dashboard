<script lang="ts">
  import { onMount, getContext } from "svelte";
  import { updateSettingValue } from "$lib/apis/settings";
  import {toast} from "svelte-sonner";

  const i18n = getContext("i18n");

  export let saveHandler: Function;
  let settings = localStorage.serverSettings ? JSON.parse(localStorage.serverSettings) : [];

  const toggleSettingUpdate = async (name, value) => {
    console.log("toggle update setting:", name, value)
    settings = await updateSettingValue(localStorage.token, name, String(value));
    localStorage.serverSettings = JSON.stringify(settings)
    toast.success(`setting ${name} updated`);
  };
</script>

<div class=" space-y-3 pr-1.5 overflow-y-scroll max-h-80">
  <div>
    {#each settings as setting}
      {#if setting.name === "allow-chat-deletion"}
        <div>
          <div class=" mb-2 text-sm font-medium">{$i18n.t("User Permissions")}</div>

          <div class="  flex w-full justify-between">
            <div class=" self-center text-xs font-medium">
              {$i18n.t("Allow Chat Deletion")}
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
        </div>
      {/if}
    {/each}
  </div>
</div>
