<script setup lang="ts">
import { ref } from "vue";
import {
	NModal,
	NForm,
	NFormItem,
	NInput,
	type FormRules,
	type FormInst,
	type FormItemRule
} from "naive-ui";
import btn from "../../components/buttons/btn.vue";
import type { CreateFileDTO } from "../../shared/types";
import modalThemes from "./modals-themes";

const show = ref<boolean>(false);
const formRef = ref<FormInst | null>(null);

const emit = defineEmits(["createFile"]);

const open = () => {
	show.value = true;
};

const close = () => {
	show.value = false;
};

defineExpose({ open, close });

const file = ref<CreateFileDTO>({
	name: ""
});
const rules: FormRules = {
	name: {
		trigger: ["input", "blur"],
		type: "string",
		validator(_rule: FormItemRule, value: string) {
			if (value.length === 0) {
				return new Error("You must provide a file name");
			}
			return true;
		}
	}
};

const createFile = () => {
	formRef.value?.validate(async errors => {
		if (!errors) emit("createFile", file.value);
	});
};
</script>

<template>
	<n-modal
		v-model:show="show"
		title="Create New File"
		:bordered="false"
		preset="card"
		style="max-width: 600px"
	>
		<n-form :model="file" :rules="rules" ref="formRef">
			<n-form-item path="name">
				<n-input
					v-model:value="file.name"
					placeholder="File Name"
					:themeOverrides="modalThemes.Input"
				/>
			</n-form-item>
		</n-form>
		<btn
			:onClick="createFile"
			text="Create"
			type="secondary"
			style="margin-top: 20px"
		/>
	</n-modal>
</template>