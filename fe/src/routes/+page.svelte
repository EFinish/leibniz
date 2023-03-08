<script lang="ts">
	import ActionBar from '../lib/components/ActionBar.svelte';
    import { subjects } from '../store';
	import Record from '$lib/components/Record.svelte';
	import { onMount } from 'svelte/types/runtime/internal/lifecycle';
	import { GetSubjects } from '$lib/utils/api';
	import type { ReadSubjectRequest } from '$lib/proto/argumentaccess/v1/server';

	onMount(async () => {
		const req: ReadSubjectRequest = {};
		const res = await GetSubjects(req)

		subjects.replace(res.subjects)
	})
</script>

<ActionBar />
<Record subjects={$subjects}/>

<style lang="postcss">
	:global(html) {
		background-color: theme(gray);
	}
</style>
