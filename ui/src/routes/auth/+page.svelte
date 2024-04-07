<script>
  import { goto } from "$app/navigation";
  import { userSignIn, userSignUp } from "$lib/apis/auths";
  import { WEBUI_API_BASE_URL, WEBUI_BASE_URL } from "$lib/constants";
  import { WEBUI_NAME, config, user } from "$lib/stores";
  import { onMount, getContext } from "svelte";
  import { toast } from "svelte-sonner";

  const i18n = getContext("i18n");

  let loaded = false;
  let mode = "signin";

  let name = "";
  let email = "";
  let password = "";

  const setSessionUser = async (sessionUser) => {
    if (sessionUser) {
      console.log("get session user:", sessionUser);
      toast.success($i18n.t(`You're now logged in.`));
      localStorage.token = sessionUser.token;
      console.log(localStorage)
      await user.set(sessionUser);
      goto("/");
    }
  };

  const signInHandler = async () => {
    const sessionUser = await userSignIn(email, password).catch((error) => {
      toast.error(error);
      return null;
    });

    await setSessionUser(sessionUser);
  };

  const signUpHandler = async () => {
    const sessionUser = await userSignUp(name, email, password).catch(
      (error) => {
        toast.error(error);
        return null;
      }
    );

    await setSessionUser(sessionUser);
  };

  const submitHandler = async () => {
    if (mode === "signin") {
      await signInHandler();
    } else {
      await signUpHandler();
    }
  };

  onMount(async () => {
    if ($user !== undefined) {
      await goto("/");
    }
    loaded = true;
  });
</script>

<svelte:head>
  <title>
    {`${$WEBUI_NAME}`}
  </title>
</svelte:head>

{#if loaded}
  <div class="fixed m-10 z-50">
    <div class="flex space-x-2">
      <div class=" self-center">
        <img
          src="{WEBUI_BASE_URL}/static/favicon.png"
          class=" w-8 rounded-full"
          alt="logo"
        />
      </div>
    </div>
  </div>

  <div
    class=" bg-white dark:bg-gray-900 min-h-screen w-full flex justify-center font-mona"
  >

    <div class="w-full sm:max-w-lg px-4 min-h-screen flex flex-col">
      <div class=" my-auto pb-10 w-full">
        <form
          class=" flex flex-col justify-center bg-white py-6 sm:py-16 px-6 sm:px-16 rounded-2xl"
          on:submit|preventDefault={() => {
            submitHandler();
          }}
        >
          <div class=" text-xl sm:text-2xl font-bold">
            {mode === "signin" ? $i18n.t("Sign in") : $i18n.t("Sign up")}
            {$i18n.t("to")}
            {$WEBUI_NAME}
          </div>

          {#if mode === "signup"}
            <div class=" mt-1 text-xs font-medium text-gray-500">
              ⓘ {$WEBUI_NAME}
              {$i18n.t(
                "does not make any external connections, and your data stays securely on your locally hosted server."
              )}
            </div>
          {/if}

          <div class="flex flex-col mt-4">
            {#if mode === "signup"}
              <div>
                <div class=" text-sm font-semibold text-left mb-1">
                  {$i18n.t("Username")}
                </div>
                <input
                  bind:value={name}
                  type="text"
                  class=" border px-4 py-2.5 rounded-2xl w-full text-sm"
                  autocomplete="name"
                  placeholder={$i18n.t("Enter Your Username")}
                  required
                />
              </div>

              <hr class=" my-3" />
            {/if}

            <div class="mb-2">
              <div class=" text-sm font-semibold text-left mb-1">
                {$i18n.t("Email")}
              </div>
              <input
                bind:value={email}
                type="email"
                class=" border px-4 py-2.5 rounded-2xl w-full text-sm"
                autocomplete="email"
                placeholder={$i18n.t("Enter Your Email")}
                required
              />
            </div>

            <div>
              <div class=" text-sm font-semibold text-left mb-1">
                {$i18n.t("Password")}
              </div>
              <input
                bind:value={password}
                type="password"
                class=" border px-4 py-2.5 rounded-2xl w-full text-sm"
                placeholder={$i18n.t("Enter Your Password")}
                autocomplete="current-password"
                required
              />
            </div>
          </div>

          <div class="mt-5">
            <button
              class=" bg-blue-900 hover:bg-blue-800 w-full rounded-full text-white font-semibold text-sm py-3 transition"
              type="submit"
            >
              {mode === "signin"
                ? $i18n.t("Sign in")
                : $i18n.t("Create Account")}
            </button>

            <div class=" mt-4 text-sm text-center">
              {mode === "signin"
                ? $i18n.t("Don't have an account?")
                : $i18n.t("Already have an account?")}

              <button
                class=" font-medium underline"
                type="button"
                on:click={() => {
                  if (mode === "signin") {
                    mode = "signup";
                  } else {
                    mode = "signin";
                  }
                }}
              >
                {mode === "signin" ? $i18n.t("Sign up") : $i18n.t("Sign in")}
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
{/if}

<style>
  .font-mona {
    font-family: "Mona Sans";
  }
</style>
