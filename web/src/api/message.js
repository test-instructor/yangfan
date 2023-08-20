import service from "@/utils/request";

const project = JSON.parse(window.localStorage.getItem("project")).ID;
const baseURL = "/message/" + project;

export const createMessage = (data) => {
  return service({
    url: baseURL + "/createMessage",
    method: "post",
    data,
  });
};

export const deleteMessage = (data) => {
  return service({
    url: baseURL + "/deleteMessage",
    method: "delete",
    data,
  });
};

export const updateMessage = (data) => {
  return service({
    url: baseURL + "/updateMessage",
    method: "put",
    data,
  });
};

export const getMessageList = (params) => {
  return service({
    url: baseURL + "/getMessageList",
    method: "get",
    params,
  });
};

export const findMessage = (params) => {
  return service({
    url: baseURL + "/findMessage",
    method: "get",
    params,
  });
};
