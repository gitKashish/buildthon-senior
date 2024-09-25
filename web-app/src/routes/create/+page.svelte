<script lang="ts">
    import { v4 as uuid } from "uuid";
    import { createPoll } from "$lib/index";
    import { blur, fade, fly } from "svelte/transition";
    import { quadInOut } from "svelte/easing";

    let loading: boolean = $state(false);
    let question: string = $state("");
    let options: string[] = $state([]);
    let locations: string[] = $state([]);
    let durationHours: number = $state(24);
    let pollId: string;

    function addPoll() {
        question = question.trim();
        if (question.length < 6) {
            alert("Question cannot be shorter than 6 character.")
            return;
        }
        if (options.length < 2) {
            alert("At least 2 options are required.");
            return;
        }
        if (locations.length === 0) {
            alert("At least 1 location is required.");
            return;
        }
        loading = true;
        pollId = uuid();
        createPoll(pollId, question, options, locations, durationHours)
            .then(() => {
                alert(
                    "Poll Created. Copy the ID to share the poll. \n" + pollId,
                );
                location.href = `/result/${pollId}`;
            })
            .catch((error) => {
                alert("Unable to create poll : " + error);
            })
            .finally(() => {
                loading = false;
            });
    }

    function addOption(event: KeyboardEvent) {
        if (event.key !== "Enter") return;
        const option: string = (event.target as HTMLInputElement).value.trim();
        if (option.length == 0) {
            return;
        }
        (event.target as HTMLInputElement).value = "";
        options.push(option);
    }

    function deleteOption(event: MouseEvent) {
        const optionIndex: string | undefined = (event.target as HTMLElement)
            .dataset.index;
        if (optionIndex === undefined) return;

        const index: number = Number(optionIndex);
        options = options.filter((_, i) => i !== index);
    }

    function addLocation(event: KeyboardEvent) {
        if (event.key !== "Enter") return;
        const location: string = (event.target as HTMLInputElement).value.trim();
        if (location.length == 0) {
            return;
        }
        (event.target as HTMLInputElement).value = "";
        locations.push(location);
    }

    function deleteLocation(event: MouseEvent) {
        const locationIndex: string | undefined = (event.target as HTMLElement)
            .dataset.index;
        if (locationIndex === undefined) return;

        const index: number = Number(locationIndex);
        locations = locations.filter((_, i) => i !== index);
    }
</script>

<div class="flex flex-col justify-center items-center min-h-screen p-6">
    <div
        class="text-center text-slate-600 text-6xl min-h-20 font-extrabold mb-6"
    >
        Create a New Poll
    </div>
    <div class="flex flex-col gap-6 w-full item max-w-2xl space-y-4">
        <div class="flex flex-col gap-4 w-full">
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                1. Poll Question
            </div>
            <textarea
                class="rounded-lg border border-gray-300 p-3 w-full focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:drop-shadow-md transition-shadow ease-in-out duration-200 disabled:bg-gray-200 disabled:text-gray-600"
                placeholder="Enter Question to ask"
                rows="6"
                required
                minlength="4"
                bind:value={question}
                disabled={loading}
            ></textarea>
        </div>

        <div class="flex flex-col gap-4">
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                2. Options
            </div>
            <div class="flex flex-col gap-2">
                {#each options as option, i}
                    <div
                        in:fade={{
                            duration: 250,
                            easing: quadInOut,
                        }}
                        out:fade={{
                            duration: 100,
                            easing: quadInOut,
                        }}
                        class="flex flex-row items-center gap-4 rounded-lg bg-indigo-400 p-2 w-full drop-shadow-md"
                    >
                        <span
                            class="text-white text-sm font-bold w-1/4 text-center"
                        >
                            Option {i + 1}
                        </span>
                        <input
                            class="rounded-lg bg-white p-3 w-full"
                            type="text"
                            disabled
                            value={option}
                            readonly
                        />
                        <button
                            class="rounded-lg bg-indigo-300 text-white font-bold text-center w-1/4 h-12 hover:bg-red-400 hover:drop-shadow-md transition duration-300 ease-in-out p-2"
                            onclick={deleteOption}
                            data-index={i}
                            disabled={loading}
                        >
                            Remove
                        </button>
                    </div>
                {/each}
            </div>
            <input
                class="rounded-lg border bg-white border-gray-300 p-3 w-full focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:drop-shadow-md transition-shadow ease-in-out duration-200 disabled:bg-gray-200 disabled:text-gray-600"
                type="text"
                placeholder="Enter new option. Press 'Enter' to add option."
                onkeydown={addOption}
                disabled={loading}
            />
        </div>

        <div class="flex flex-col gap-4">
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                3. Locations
            </div>
            <div class="flex flex-wrap gap-2 justify-center">
                {#each locations as location, i}
                    <div
                        in:fade={{
                            duration: 250,
                            easing: quadInOut,
                        }}
                        out:fade={{
                            duration: 100,
                            easing: quadInOut,
                        }}
                        class="flex flex-row items-center gap-4 rounded-lg bg-indigo-400 p-2 drop-shadow-md w-fit"
                    >
                        <input
                            class="rounded-lg bg-white p-3 w-full"
                            type="text"
                            disabled
                            value={location}
                            readonly
                        />
                        <button
                            class="rounded-lg bg-indigo-300 text-white font-bold text-center w-min h-12 hover:bg-red-400 hover:drop-shadow-md transition duration-300 ease-in-out p-2 disabled:bg-indigo-200"
                            onclick={deleteLocation}
                            data-index={i}
                            disabled={loading}
                        >
                            Remove
                        </button>
                    </div>
                {/each}
            </div>
            <input
                class="rounded-lg border bg-white border-gray-300 p-3 w-full focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:drop-shadow-md transition-shadow ease-in-out duration-200 disabled:bg-gray-200 disabled:text-gray-600"
                type="text"
                placeholder="Enter new Location. Press 'Enter' to add location."
                onkeydown={addLocation}
                disabled={loading}
            />
        </div>

        <div class="flex flex-col gap-4 w-full">
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                4. Poll Duration (in Hours)
            </div>
            <input
                class="rounded-lg border bg-white border-gray-300 p-3 w-full focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:drop-shadow-md transition-shadow ease-in-out duration-200 disabled:bg-gray-200 disabled:text-gray-600"
                type="number"
                placeholder="Enter duration in hours"
                bind:value={durationHours}
                min="1"
                disabled={loading}
            />
        </div>
        <button
            class="rounded-lg bg-indigo-500 text-white w-1/2 mx-auto font-medium h-12 hover:bg-indigo-600 hover:drop-shadow-md transition duration-200 ease-in-out disabled:bg-indigo-200 disabled:text-indigo-500 disabled:border-0"
            onclick={addPoll}
            disabled={loading}
        >
            {loading ? "Loading" : "Submit"}
        </button>
    </div>
</div>
