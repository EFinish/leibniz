<script lang="ts">
	import type { Subject } from '$lib/proto/argumentaccess/v1/argument';
	import { Li, List } from 'flowbite-svelte';
	import SubjectRecord from './SubjectRecord.svelte';

	export let subjects: Subject[] = [];
</script>

<SubjectRecord subjects={subjects} />