import { WEBUI_API_BASE_URL } from "$lib/constants";
import {toast} from "svelte-sonner";

const setupSettings = async () => {
  let settings = await getAllSettings(localStorage.token);
  localStorage.serverSettings = JSON.stringify(settings)
};

export const getAllSettings = async (token: string) => {
  let error = null;

  const res = await fetch(`${WEBUI_API_BASE_URL}/settings/`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  })
    .then(async (res) => {
      if (!res.ok) throw await res.json();
      return res.json();
    })
    .catch((err) => {
      console.log(err);
      error = err;
      return null;
    });

  if (error) {
    throw error;
  }

  return res;
};

export const updateSettingValue = async (
    token: string,
    name: string,
    value: string
) => {
  let error = null;

  const res = await fetch(`${WEBUI_API_BASE_URL}/settings/`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      name: name,
      value: value,
    }),
  })
      .then(async (res) => {
        if (!res.ok) throw await res.json();
        return res.json();
      })
      .catch((err) => {
        console.log(err);
        return err;
      });

  if (error) {
    throw error;
  }

  return res;
};

setupSettings();
