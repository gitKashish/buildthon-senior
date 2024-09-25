<script lang="ts">
    import { page } from "$app/stores";
    import type { PollResponse, PollDetails } from "$lib/ambient";
    import { submitVote, getPoll, isPollOpen } from "$lib/index";
    import { quadInOut } from "svelte/easing";
    import { fly } from "svelte/transition";

    const pollId = $page.params.pollId;

    let loading: boolean = $state(false);
    let question: string = $state("");
    let options: string[] = $state([]);
    let locations: string[] = $state([]);

    let selectedOption: string = $state("");
    let selectedLocation: string = $state("");

    let pollDetails: PollDetails;

    $effect(() => {
        if (pollId.length != 36) {
            alert("Invalid Poll ID. Returning to home.");
            location.href = `/`;
        }
        loading = true;
        isPollOpen(pollId)
            .then((obj) => {
                if (obj.result.result) {
                    getPoll(pollId)
                        .then((obj) => {
                            let pollResponse: PollResponse = obj;
                            pollDetails = pollResponse.result.result;
                            question = pollDetails.question;
                            options = pollDetails.options;
                            locations = pollDetails.locations;
                        })
                        .catch((error) => {
                            window.alert(
                                "Unable to fetch poll details : " + error,
                            );
                        });
                } else {
                    location.href = `/result/${pollId}`;
                }
            })
            .catch((error) => {
                alert("Invalid Poll Id : " + error);
            })
            .finally(() => {
                loading = false;
            });
    });

    function selectOption(option: string) {
        selectedOption = option;
    }

    function selectLocation(location: string) {
        selectedLocation = location;
    }

    function addVote() {
        if (!selectedOption || !selectedLocation) {
            window.alert("Please select both an option and a location.");
            return;
        }
        loading = true;
        let optionIndex: number = options.indexOf(selectedOption);
        let locationIndex: number = locations.indexOf(selectedLocation);
        submitVote(pollId, optionIndex, locationIndex)
            .then(() => {
                location.href = "/result/" + pollId;
            })
            .catch((error) => {
                alert("Unable to submit vote : " + error);
            })
            .finally(() => {
                loading = false;
            });
    }
</script>

{#if loading && options.length == 0 && locations.length == 0}
    <span
        class="flex flex-col justify-center text-center text-slate-600 text-6xl font-extrabold min-h-screen"
        >Loading...</span
    >
{:else}
    <div
        class="flex flex-col justify-center items-center min-h-screen max-w-lg m-auto p-6"
    >
        <div class="text-slate-600 text-4xl mb-2 font-extrabold w-full">
            Vote : {question}
        </div>
        <span class="text-left text-slate-400 text-lg font-bold w-full mb-8">
            Poll ID : {pollId}
        </span>
        <div class="flex flex-col gap-6 w-full space-y-4">
            <div class="flex flex-col gap-4">
                <div
                    class="flex items-left text-slate-500 text-xl font-extrabold"
                >
                    1. Choose an Option
                </div>
                <div class="flex flex-col gap-2">
                    {#each options as option, i}
                        <button
                            in:fly={{
                                x: -200,
                                duration: 300,
                                delay: 0 + i * 30,
                                easing: quadInOut,
                            }}
                            class="{selectedOption == option
                                ? 'bg-indigo-500 text-white drop-shadow-lg hover:bg-indigo-500'
                                : 'bg-white text-gray-700 hover:bg-indigo-400'} border-2 border-indigo-500 w-full p-4 text-center rounded-lg font-bold transition duration-300 ease-in-out hover:drop-shadow-lg hover:text-white"
                            onclick={() => selectOption(option)}
                        >
                            {option}
                        </button>
                    {/each}
                </div>
            </div>

            <div class="flex flex-col gap-4">
                <div
                    class="flex items-left text-slate-500 text-xl font-extrabold"
                >
                    2. Choose Your Location
                </div>
                <div class="flex flex-col gap-2">
                    {#each locations as location, i}
                        <button
                            in:fly={{
                                x: -200,
                                duration: 300,
                                delay: 0 + i * 30,
                                easing: quadInOut,
                            }}
                            class="{selectedLocation == location
                                ? 'bg-indigo-500 text-white drop-shadow-lg hover:bg-indigo-500'
                                : 'bg-white text-gray-700 hover:bg-indigo-400'} border-2 border-indigo-500 w-full p-4 text-center rounded-lg font-bold transition duration-300 ease-in-out hover:drop-shadow-lg hover:text-white"
                            onclick={() => selectLocation(location)}
                        >
                            {location}
                        </button>
                    {/each}
                </div>
            </div>

            <button
                class="rounded-lg bg-indigo-500 text-white font-medium w-full h-12 hover:bg-indigo-600 transition duration-300 ease-in-out disabled:bg-indigo-200 disabled:text-indigo-500"
                onclick={addVote}
                disabled={loading}
            >
                {loading ? "Loading..." : "Submit Vote"}
            </button>
        </div>
    </div>
{/if}
