<template>
    <div>
        <div v-if="loading">Loading...</div>
        <div v-else-if="error">{{ error }}</div>
        <div v-else v-for="item in repos" v-bind:key="item.id">{{ item.username }}/{{ item.name }}</div>
    </div>
</template>

<script lang="ts">
import { RepositoryInfo } from '@/model/Repository';
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'RepoListing',
    components: {
    },
    data() {
        return {
            loading: true,
            repos: null as Array<RepositoryInfo> | null,
            error: null as string | null,
        }
    },
    created() {
        // watch the params of the route to fetch the data again
        this.$watch(
            () => this.$route.params,
            () => {
                this.fetchRepos()
            },
            // fetch the data when the view is created and the data is
            // already being observed
            { immediate: true }
        )
    },
    methods: {
        async fetchRepos() {
            this.error = this.repos = null
            this.loading = true

            try {
                let response = await (await fetch("/api/v1/repos")).json();
                this.repos = response as Array<RepositoryInfo>;
            } catch (e: unknown) {
                this.error = String(e)
            } finally {
                this.loading = false
            }
        },
    }
});
</script>
