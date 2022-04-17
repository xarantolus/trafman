<script lang="ts">
import { defineComponent, type PropType } from 'vue'
import type { RepositoryInfo } from '@/model/Repository';

export default defineComponent({
    name: "repository-card",
    props: {
        repo: {
            type: Object as PropType<RepositoryInfo>,
            required: true
        }
    },
    components: {},
    methods: {}
})
</script>

<template>
    <div class="card repo-box">
        <div class="card-content">
            <span v-if="repo.stars > 0" class="card-header-icon stars">⭐{{ repo.stars }}</span>
            <span v-if="repo.is_fork" class="card-header-icon fork">🍴</span>
            <h4 class="title is-4">
                <!-- TODO: Add router link to some kind of overview page for a repository -->
                <a class="repocard-link" :href="'https://github.com/' + repo.username + '/' + repo.name">{{ repo.username }}/{{ repo.name }}</a>
            </h4>
            <p v-if="repo.description.trim()" class="content">{{repo.description}}</p>
        </div>
    </div>
</template>

<style>
.repocard-link::after {
    content: '';
    position: absolute;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
}

.repo-box {
    background: var(--card-color);
    border: 3px solid var(--border-color);
    margin-bottom: 2%;
}

.stars {
    position: absolute;
    top: 0;
    right: 0;
    user-select: none;
}

.fork {
    position: absolute;
    bottom: 0;
    right: 0;
    user-select: none;
}
</style>
