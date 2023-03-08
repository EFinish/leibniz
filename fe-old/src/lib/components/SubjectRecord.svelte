<script lang="ts">
	import type { Subject } from '$lib/proto/argumentaccess/v1/argument';
	import { Li, List } from 'flowbite-svelte';

    export let subjects: Subject[] = [];
</script>

<List>
    {#each subjects as subject (subject.id)}
        <Li>{subject.body}</Li>
    {/each}
</List>