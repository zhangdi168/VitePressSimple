import { createApp } from "vue";
import { createPinia } from "pinia";
import "./style.scss";
import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import {
  Alert,
  Button,
  Divider,
  Dropdown,
  Empty,
  Form,
  Input,
  InputNumber,
  List,
  Menu,
  message,
  Modal,
  Popconfirm,
  Radio,
  Switch,
  Tabs,
  Tag,
  Tooltip,
  Tree,
} from "ant-design-vue";

// Import icon libraries
import { Quasar } from "quasar";
import "@quasar/extras/material-icons/material-icons.css";
// Import Quasar css
import "quasar/src/css/index.sass";
import { InstallCodemirro } from "codemirror-editor-vue3";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(i18n);
app.use(Quasar, {
  plugins: {}, // import Quasar plugins and add here
});
//局部注册Antd
app
  .use(Button)
  .use(Tree)
  .use(Input)
  .use(Form)
  .use(Tabs)
  .use(List)
  .use(Switch)
  .use(InputNumber)
  .use(Radio)
  .use(Modal)
  .use(Dropdown)
  .use(Menu)
  .use(Empty)
  .use(Alert)
  .use(Tag)
  .use(Divider)
  .use(Popconfirm)
  .use(Tooltip);

app.use(InstallCodemirro);

app.config.globalProperties.$message = message;

app.mount("#app");
