<template>
    <Line :chart-options="(chartOptions as any)" :chart-data="data" />
</template>

<script lang="ts">
import { Line } from 'vue-chartjs';
import 'chartjs-adapter-luxon';
import { DateTime } from 'luxon';
import 'chart.js/auto';
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'simple-date-chart',
    props: {
        chartTitle: {
            type: String,
            required: true
        },
        data: {
            type: Object as any,
            required: true
        }
    },
    components: {
        Line,
    },
    data() {
        // eslint-disable-next-line @typescript-eslint/no-this-alias
        let vueRef = this;
        const dayFormat = 'DD T';
        return {
            chartOptions: {
                maintainAspectRatio: false,
                plugins: {
                    title: {
                        text: 'Clones over time',
                        display: true
                    },
                    tooltip: {
                        callbacks: {
                            title: function (arg: Array<any>) {
                                let idx = arg[0].dataIndex;
                                let data = vueRef.$props.data.labels[idx];
                                return DateTime.fromISO(data).toLocaleString({ month: 'long', day: 'numeric' });
                            }
                        }
                    }
                },
                interaction: {
                    intersect: false,
                    mode: 'index'
                },
                tooltips: {
                    mode: 'index',
                    intersect: false
                },
                hover: {
                    mode: 'index',
                    intersect: false
                },
                scales: {
                    x: {
                        type: 'time',
                        time: {
                            // Luxon format string
                            tooltipFormat: dayFormat,
                            unit: 'day'
                        },
                        display: true,
                        title: {
                            display: true,
                            text: 'Date',
                        }
                    },
                    y: {
                        title: {
                            display: true,
                            text: 'Count'
                        }
                    }
                },
            }
        }
    }
});
</script>
