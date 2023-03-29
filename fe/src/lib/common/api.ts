export function apiPut<T>(path: string, data: T): Promise<Response> {
  return fetch(path, {
    method: "PUT",
    headers: {
      "content-type": "application/json",
    },
    body: JSON.stringify(data),
  });
}

export function apiPost<T>(path: string, data: T): Promise<Response> {
  return fetch(path, {
    method: "POST",
    headers: {
      "content-type": "application/json",
    },
    body: JSON.stringify(data),
  });
}
export function apiDelete<T>(path: string, data: T): Promise<Response> {
  return fetch(path, {
    method: "DELETE",
    headers: {
      "content-type": "application/json",
    },
    body: JSON.stringify(data),
  });
}

export function apiGet(path: string): Promise<Response> {
  return fetch(path, {
    method: "GET",
    headers: {
      "content-type": "application/json",
    },
  });
}
