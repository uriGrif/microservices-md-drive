<script setup lang="ts">
import Layout from "../components/layout.vue";
import {
	NForm,
	NInput,
	NFormItem,
	NSelect,
	NConfigProvider,
	useMessage,
	type FormInst,
	type FormRules,
	type FormItemRule
} from "naive-ui";
import btn from "../components/buttons/btn.vue";
import { ref, toRaw } from "vue";
import { codes } from "iso-country-codes";
import formThemeOverride from "../shared/formThemeOverride";
import { updateProfile } from "../services/profiles";
import { ApiError } from "../services/api";
import { useAuth0 } from "@auth0/auth0-vue";
import { profileStore } from "../stores/profile";
import router from "../router";

const { getAccessTokenSilently, user } = useAuth0();

const message = useMessage();
const formRef = ref<FormInst | null>(null);
const formValue = ref({
	nickname: profileStore.profile?.nickname || null,
	country: profileStore.profile?.country || null,
	description: profileStore.profile?.description || null
});
const rules: FormRules = {
	nickname: [
		{
			required: true,
			message: "You must provide a nickname",
			trigger: ["input", "blur"]
		},
		{
			trigger: ["input", "blur"],
			type: "string",
			validator(_rule: FormItemRule, value: string) {
				if (value.length > 0 && !/^[A-Za-z0-9]+$/.test(value)) {
					return new Error(
						"Nickname can only contain upper or lower case letters and numbers"
					);
				}
				return true;
			}
		}
	],
	country: {
		required: true,
		message: "You must select a country",
		trigger: "blur"
	}
};

const countries = (codes as Array<{ name: string; alpha2: string }>).map(c => {
	return {
		label: c.name,
		value: c.alpha2
	};
});

const handleValidateClick = () => {
	formRef.value?.validate(async errors => {
		if (!errors) {
			try {
				const token = await getAccessTokenSilently();
				const profile = await updateProfile(
					toRaw(formValue.value),
					user.value?.sub || "",
					token
				);
				profileStore.setProfile(profile);
				message.success("Profile successfully updated!");
				router.push("/update-profile");
			} catch (error) {
				if (error instanceof ApiError) {
					message.error(error.message);
				}
			}
		}
	});
};
</script>

<template>
	<Layout>
		<h1>Update your profile</h1>
		<n-config-provider :theme-overrides="formThemeOverride">
			<n-form ref="formRef" :model="formValue" :rules="rules">
				<n-form-item label="Nickname" path="nickname">
					<n-input
						v-model:value="formValue.nickname"
						placeholder="Input Nickname"
					/>
				</n-form-item>
				<n-form-item label="Country" path="country">
					<n-select
						v-model:value="formValue.country"
						placeholder="Select a country"
						:options="countries"
						filterable
					/>
				</n-form-item>
				<n-form-item label="Description" path="description">
					<n-input
						v-model:value="formValue.description"
						placeholder="Write a description about yourself"
						type="textarea"
					/>
				</n-form-item>
			</n-form>
		</n-config-provider>
		<btn :onClick="handleValidateClick" text="Submit" />
	</Layout>
</template>

<style scoped>
.btn {
	float: right;
}
</style>
