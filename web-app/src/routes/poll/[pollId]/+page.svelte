<script lang="ts">
    import { page } from "$app/stores";
    import type { PollResponse, PollDetails } from "$lib/ambient";
    import { submitVote, getPoll, isPollOpen } from "$lib/index";

    const pollId = $page.params.pollId;

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
            });
    });

    function selectOption(option: string) {
        selectedOption = option;
    }

    function selectLocation(location: string) {
        selectedLocation = location;
    }

    async function addVote() {
        if (!selectedOption || !selectedLocation) {
            window.alert("Please select both an option and a location.");
            return;
        }
        try {
            let optionIndex: number = options.indexOf(selectedOption);
            let locationIndex: number = locations.indexOf(selectedLocation);
            await submitVote(pollId, optionIndex, locationIndex);
            alert("Vote submitted succesfully.");
            location.href = "/result/" + pollId;
        } catch (error) {
            window.alert("Error submitting your vote.");
        }
    }
</script>

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
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                1. Choose an Option
            </div>
            <div class="flex flex-col gap-2">
                {#each options as option}
                    <button
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
            <div class="flex items-left text-slate-500 text-xl font-extrabold">
                2. Choose Your Location
            </div>
            <div class="flex flex-col gap-2">
                {#each locations as location, i}
                    <button
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
            class="rounded-lg border border-indigo-500 text-indigo-500 font-medium w-full h-12 hover:bg-indigo-500 hover:text-white transition duration-300 ease-in-out"
            onclick={addVote}
        >
            Submit Vote
        </button>
    </div>
</div>
