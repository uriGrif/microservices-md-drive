<script setup lang="ts">
interface Props {
	text: String;
	type?: "primary-light" | "primary-dark" | "secondary";
	onClick?: () => void;
	href?: String;
}
const props = withDefaults(defineProps<Props>(), {
	type: "primary-light"
});

const handleClick = () => {
	props.onClick?.();
	if (props.href) {
	}
};
</script>

<template>
	<button @click="handleClick" :class="`btn ${type}-btn`">{{ text }}</button>
</template>

<style>
.btn {
	border-radius: var(--border-radius);
	border: 2px solid var(--primary-light-color);
	padding: 0.8rem 2rem;
	font-family: var(--font);
	font-size: 1rem;
	transition: 100ms;
}

.btn:hover {
	cursor: pointer;
}

.primary-light-btn {
	background-color: var(--primary-light-color);
}

.primary-dark-btn {
	background-color: transparent;
	color: var(--primary-light-color);
}
</style>
