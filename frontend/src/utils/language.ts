import { useI18n } from "vue-i18n";

//封装语言包读取函数，方便快速调用
export const lang = (k: string) => {
  const { t } = useI18n();
  return t(k);
};
