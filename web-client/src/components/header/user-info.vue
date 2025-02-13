<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";
import { NDropdown, type GlobalThemeOverrides } from "naive-ui";
import { NConfigProvider } from "naive-ui";
import { profileStore } from "../../stores/profile";
import router from "../../router";

const { user, logout } = useAuth0();

const handleLogout = () => {
	profileStore.clearProfile();
	logout({
		logoutParams: {
			returnTo: window.location.origin
		}
	});
};

const dropdownOptions = [
	{
		label: "Update Profile",
		key: "update"
	},
	{
		label: "Log Out",
		key: "logout"
	}
];

const dropdownTheme: GlobalThemeOverrides = {
	Dropdown: {
		color: "var(--primary-light-color)", // Background color
		textColor: "var(--primary-dark-color)", // Text color
		optionColorHover: "var(--secondary-color)", // Hover background color
		optionTextColorHover: "var(--text-color)", // Hover background color
		optionTextColor: "var(--primary-dark-color)" // Option text color
	}
};

const handleSelect = (key: string | number) => {
	switch (key) {
		case "logout":
			handleLogout();
			break;
		case "update":
			router.push("/update-profile");
			break;
	}
};
</script>

<template>
	<n-config-provider :theme-overrides="dropdownTheme">
		<n-dropdown
			trigger="click"
			:options="dropdownOptions"
			@select="handleSelect"
			class="userInfoDropdown"
		>
			<div class="userInfo">
				<img
					class="userPicture"
					:src="user?.picture"
					alt="profile picture"
				/>
				<!-- TODO usar si existe picture del profile -->
				<div class="userInfoNames">
					<span class="userInfoName">{{ user?.name }}</span>
					<span>
						@{{ profileStore.profile?.nickname || "<>" }}
						<!-- TODO cambiar por nickname del profile -->
					</span>
				</div>
			</div>
		</n-dropdown>
	</n-config-provider>
</template>

<style scoped>
.userInfo {
	border-radius: 5px;
	border: 2px solid var(--primary-light-color);
	color: var(--primary-dark-color);
	padding: 0.4rem 0.8rem;
	font-family: var(--font);
	font-size: 1rem;
	margin: 5px;
	transition: 100ms;
	background-color: var(--primary-light-color);
	display: flex;
	align-items: center;
}

.userInfo:hover {
	cursor: pointer;
}

.userPicture {
	border-radius: 50%;
	height: 32px;
	width: 32px;
	margin-right: 20px;
}

.userInfoNames {
	display: flex;
	flex-direction: column;
}

.userInfoName {
	font-weight: 600;
}
</style>
