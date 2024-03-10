import { createI18n } from "vue-i18n";
import { zhHans } from "@/i18n/locales/zh-Hans";
import { en } from "@/i18n/locales/en";

const i18n = createI18n({
  locale: "zh-Hans",
  fallbackLocale: "en",
  legacy: false,
  messages: {
    "zh-Hans": zhHans,
    en: en,
  },
});

export default i18n;
