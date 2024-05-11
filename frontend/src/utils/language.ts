import { useI18n } from "vue-i18n";
import { VitePressVersion } from "@/configs/cnts";

//封装语言包读取函数，方便快速调用
export const lang = (k: string) => {
  if (k == "vitePressVersion") {
    return VitePressVersion;
  }
  const { t } = useI18n();
  return t(k);
};
