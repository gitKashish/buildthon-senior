<script lang="ts">
    import { page } from "$app/stores";
    import { isPollOpen } from "$lib/index";
    const pollId = $page.params.pollId;

    let loadingText = $state("");

    $effect(() => {
        if (pollId.length != 36) {
            loadingText = "Invalid Poll ID. Redirecting to Home...";
            location.href = `/`;
        }
        loadingText = "Checking Poll status..."
        isPollOpen(pollId)
            .then((obj) => {
                if (obj.result.result) {
                    location.href = `/poll/${pollId}`;
                } else {
                    location.href = `/result/${pollId}`;
                }
            })
            .catch(() => {
                loadingText = "Poll ID does not exist. Redirecting to Home...";
                location.href = `/`;
            });
    });
</script>

<span
    class="flex flex-col justify-center text-center text-slate-600 text-4xl font-extrabold min-h-screen"
>
    {loadingText}
</span>
