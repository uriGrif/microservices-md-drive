<script setup lang="ts">
import btn from "../buttons/btn.vue";
import userInfo from "./user-info.vue";
import { RouterLink } from "vue-router";
import { Home12Regular } from "@vicons/fluent";
import { NIcon } from "naive-ui";
import { useAuth0 } from "@auth0/auth0-vue";

const { loginWithRedirect, isAuthenticated, isLoading } = useAuth0();

const handleLogin = () => {
	loginWithRedirect();
};

const handleSignUp = () => {
	loginWithRedirect({
		authorizationParams: {
			screen_hint: "signup"
		}
	});
};
</script>

<template>
	<div class="my-header">
		<RouterLink to="/">
			<n-icon
				:component="Home12Regular"
				size="2rem"
				color="var(--primary-light-color)"
			/>
		</RouterLink>
		<div v-if="isLoading">Loading</div>
		<template v-else>
			<div v-if="!isAuthenticated">
				<btn
					text="Sign Up"
					type="primary-dark"
					:on-click="handleSignUp"
				/>
				<btn text="Log In" :on-click="handleLogin" />
			</div>
			<div v-else>
				<userInfo />
			</div>
		</template>
	</div>
</template>

<style>
.my-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 20px 2rem;
}

.my-header .btn {
	margin: 0 5px;
}
</style>
