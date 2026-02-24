import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import auth0Client from "./shared/auth0-client";

createApp(App)
	.use(router)
	.use(auth0Client)
	.mount("#app");
