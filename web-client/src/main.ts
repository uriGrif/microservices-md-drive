import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { createAuth0 } from "@auth0/auth0-vue";
import router from "./router";

createApp(App)
	.use(router)
	.use(
		createAuth0({
			domain: import.meta.env.VITE_AUTH0_DOMAIN,
			clientId: import.meta.env.VITE_AUTH0_CLIENT_ID,
			authorizationParams: {
				audience: import.meta.env.VITE_AUTH0_AUDIENCE,
				redirect_uri: import.meta.env.VITE_AUTH0_CALLBACK_URL
			}
		})
	)
	.mount("#app");
