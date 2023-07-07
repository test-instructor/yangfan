import service from "@/utils/request";

const project = JSON.parse(window.localStorage.getItem("project")).ID;
const baseURL = "/performance/" + project;

export const getPerformanceList = (params) => {
  return service({
    url: baseURL + "/getPerformanceList",
    method: "get",
    params,
  });
};

export const createPerformance = (data, params) => {
  return service({
    url: baseURL + "/createPerformance",
    method: "post",
    data,
    params,
  });
};

export const addOperation = (data) => {
  return service({
    url: baseURL + "/addOperation",
    method: "post",
    data,
  });
};

export const deletePerformance = (data, params) => {
  return service({
    url: baseURL + "/deletePerformance",
    method: "delete",
    data,
    params,
  });
};

export const deleteReport = (data) => {
  console.log("=============", data);
  return service({
    url: baseURL + "/deleteReport",
    method: "delete",
    data,
  });
};

export const deletePerformanceIds = (data, params) => {
  return service({
    url: baseURL + "/deletePerformanceIds",
    method: "delete",
    data,
    params,
  });
};

export const findPerformance = (params) => {
  return service({
    url: baseURL + "/findPerformance",
    method: "get",
    params,
  });
};

export const updatePerformance = (data, params) => {
  return service({
    url: baseURL + "/updatePerformance",
    method: "put",
    data,
    params,
  });
};

export const addPerformanceCase = (data) => {
  return service({
    url: baseURL + "/addPerformanceCase",
    method: "post",
    data,
  });
};

export const findPerformanceCase = (params) => {
  return service({
    url: baseURL + "/findPerformanceCase",
    method: "get",
    params,
  });
};

export const setPerformanceCase = (data) => {
  return service({
    url: baseURL + "/setPerformanceCase",
    method: "post",
    data: data,
  });
};

export const sortPerformanceCase = (data) => {
  return service({
    url: baseURL + "/sortPerformanceCase",
    method: "post",
    data,
  });
};

export const delPerformanceCase = (data) => {
  return service({
    url: baseURL + "/delPerformanceCase",
    method: "delete",
    data,
  });
};

export const findPerformanceStep = (params) => {
  return service({
    url: baseURL + "/findPerformanceStep",
    method: "get",
    params,
  });
};

export const getReportList = (params) => {
  return service({
    url: baseURL + "/getReportList",
    method: "get",
    params,
  });
};

export const findReport = (params) => {
  return service({
    url: baseURL + "/findReport",
    method: "get",
    params,
  });
};
