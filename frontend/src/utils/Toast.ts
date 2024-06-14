import { message } from "ant-design-vue";
import { IsEmptyValue } from "@/utils/utils";

export function ToastInfo(msg: string) {
  message.info(msg);
}

export function ToastError(msg: string) {
  if (IsEmptyValue(msg.trim())) return;
  message.error(msg);
}

export function ToastSuccess(msg: string) {
  if (IsEmptyValue(msg.trim())) return;
  message.success(msg);
}

//根据message是否为空字符串判断是否成功
export const ToastCheck = (
  message: string,
  successText: string = "操作成功",
) => {
  if (message == "" && successText != "") {
    ToastSuccess(successText);
    return true;
  } else {
    ToastError(message);
    return false;
  }
};
