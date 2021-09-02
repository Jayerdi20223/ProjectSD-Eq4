import { API_HOST } from "../utils/constant";
import { getTokenApi } from "./auth";

export function getUserApi(id) {
  const url = `${API_HOST}/verperfil?id=${id}`;

  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${getTokenApi()}`
    }
  };

  return fetch(url, params)
    .then(response => {
      // eslint-disable-next-line no-throw-literal
      if (response.status >= 400) throw null;
      return response.json();
    })
    .then(result => {
      return result;
    })
    .catch(err => {
      return err;
    });
}

export function uploadBannerApi(file) {
  const url = `${API_HOST}/subirBanner`;

  const formData = new FormData();
  formData.append("banner", file);

  const params = {
    method: "POST",
    headers: {
      Authorization: `Bearer ${getTokenApi()}`
    },
    body: formData
  };

  return fetch(url, params)
    .then(response => {
      return response.json(); /*llega el header*/
    })
    .then(result => {
      return result; /*llega el resultado decodificado*/
    })
    .catch(err => {
      return err;
    });
}

export function uploadAvatarApi(file) { /*Función q se conecta al endpoint*/
  const url = `${API_HOST}/subirAvatar`;

  const formData = new FormData();
  formData.append("avatar", file);

  const params = {
    method: "POST",
    headers: {
      Authorization: `Bearer ${getTokenApi()}`
    },
    body: formData
  };

  return fetch(url, params)
    .then(response => {
      return response.json();
    })
    .then(result => {
      return result;
    })
    .catch(err => {
      return err;
    });
}

export function updateInfoApi(data) {
  const url = `${API_HOST}/modificarPerfil`;

  const params = {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${getTokenApi()}`
    },
    body: JSON.stringify(data) /*data de la info del user*/
  };

  return fetch(url, params)
    .then(response => {
      return response;
    })
    .catch(err => {
      return err;
    });
}

export function followUserApi(idUser) {
  const url = `${API_HOST}/altaRelacion?id=${idUser}`;

  const params = {
    method: "POST",
    headers: {
      Authorization: `Bearer ${getTokenApi()}`,
    },
  };

  return fetch(url, params)
    .then((response) => {
      return response.json();
    })
    .then((result) => {
      return result;
    })
    .catch((err) => {
      return err;
    });
}
