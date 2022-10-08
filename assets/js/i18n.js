import { createI18n } from "vue-i18n";
<<<<<<< HEAD:assets/js/i18n.js
import de from "../i18n/de.toml";
import en from "../i18n/en.toml";
import it from "../i18n/it.toml";
import lt from "../i18n/lt.toml";
=======
import de from "./de.toml";
import en from "./en.toml";
import it from "./it.toml";
import lt from "./lt.toml";
>>>>>>> bfd983beb (change translation format to form js to toml):assets/js/i18n/index.js

const PREFERRED_LOCALE_KEY = "preferred_locale";

function getBrowserLocale() {
  const navigatorLocale =
    navigator.languages !== undefined ? navigator.languages[0] : navigator.language;
  if (!navigatorLocale) {
    return undefined;
  }
  return navigatorLocale.trim().split(/-|_/)[0];
}

export default createI18n({
  locale: window.localStorage[PREFERRED_LOCALE_KEY] || getBrowserLocale(),
  fallbackLocale: "en",
  messages: { de, en, it, lt },
});
