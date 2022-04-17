<template>
    <div>
        <div v-if="loading">Loading...</div>
        <div v-else-if="error">{{ error }}</div>
        <repository-card :repo="stats!.repository" />
    </div>
</template>

<script lang="ts">
import RepositoryCard from '@/components/RepositoryCard.vue';
import { defineComponent } from 'vue';
import { RepoStats } from '@/model/RepoStats';

export default defineComponent({
    name: 'repo-page',
    components: {
        RepositoryCard,
    },
    data() {
        return {
            loading: true,
            stats: null as RepoStats | null,
            error: null as string | null,
        }
    },
    created() {
        // watch the params of the route to fetch the data again
        this.$watch(
            () => this.$route.params,
            (toParams: any) => {
                this.fetchRepoStats(toParams.username, toParams.reponame)
            },
            // fetch the data when the view is created and the data is
            // already being observed
            { immediate: true }
        )
    },
    methods: {
        async fetchRepoStats(username: string, reponame: string) {
            this.error = this.stats = null
            this.loading = true

            try {
                let response = await (await fetch(`/api/v1/repo/${username}/${reponame}/stats`)).json();
                this.stats = response as RepoStats;
            } catch (e: unknown) {
                this.error = String(e)
            } finally {
                this.loading = false
            }
        },
    }
});
</script>
