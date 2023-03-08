import type { Subject } from '$lib/proto/argumentaccess/v1/argument';
import type { ReadSubjectRequest } from '$lib/proto/argumentaccess/v1/server';
import { GetSubjects } from '$lib/utils/api';
import { writable } from 'svelte/store';

async function CreateSubjects() {
    // const req: ReadSubjectRequest = {};
    // const res = await GetSubjects(req);
	// console.log(res);
    // const subjectsInit: Subject[] = res.subjects
	const subjectsInit: Subject[]  = [];
	const { subscribe, set } = writable(subjectsInit);

	return {
		subscribe,
        // TODO
		replace: (subjects:Subject[]) => set(subjects)
		// add: (potato: Potato) => update(n => n + 1),

		// remove: (potato: Potato) => update(n => n - 1),
	};
}

export const subjects = await CreateSubjects();