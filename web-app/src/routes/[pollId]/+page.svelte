<script lang="ts">
    import { page } from "$app/stores";
    import { isPollOpen } from "$lib/index";
    const pollId = $page.params.pollId;

    $effect(() => {
        if (pollId.length != 36) {
            alert("Invalid Poll ID. Returning to home.");
            location.href = `/`;
        }
        isPollOpen(pollId)
            .then((obj) => {
                if (obj.result.result) {
                    location.href = `/poll/${pollId}`;
                } else {
                    location.href = `/result/${pollId}`;
                }
            })
            .catch((error) => {
                alert("Invalid Poll Id : " + error);
            });
    });
</script>
