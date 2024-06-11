import { VitePressVersion } from "@/configs/cnts";
import i18n from "@/i18n";
//封装语言包读取函数，方便快速调用
export const lang = (k: string) => {
  if (k == "vitePressVersion") {
    return VitePressVersion;
  }
  const { t } = i18n.global;
  return t(k);
};
