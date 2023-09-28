import service from "@/utils/request";

const baseURL = "/ci/resp/";

export const findCIReport = (params) => {
  // ci
  return service({
    url: baseURL + "/findReport",
    method: "get",
    params,
  });
};

export const getCIReportDetail = (params) => {
  // ci
  return service({
    url: baseURL + "/getReportDetail",
    method: "get",
    params,
    donNotShowLoading: true,
  });
};
// http://localhost:8080/#/response/1/1/1
