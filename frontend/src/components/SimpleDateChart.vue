<template>
    <Line :chart-options="(chartOptions as any)" :chart-data="data" />
</template>

<script lang="ts">
import { Line } from 'vue-chartjs';
import 'chartjs-adapter-luxon';
import { DateTime } from 'luxon';
import 'chart.js/auto';
import { defineComponent } from 'vue';
import { isDarkMode } from '@/mixins/dark_mode';

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
        const gridColor = isDarkMode() ? "#aaa" : "#333";
        const titleColor = isDarkMode() ? "#fff" : "#333";
        const ticksColor = isDarkMode() ? "#ddd" : "#333";
        const labelColor = isDarkMode() ? "#eee" : "#222";

        return {
            chartOptions: {
                maintainAspectRatio: false,
                plugins: {
                    title: {
                        text: this.$props.chartTitle,
                        display: true,
                        color: titleColor
                    },
                    tooltip: {
                        callbacks: {
                            title: function (arg: Array<any>) {
                                let idx = arg[0].dataIndex;
                                let data = vueRef.$props.data.labels[idx];
                                return DateTime.fromISO(data).toLocaleString({ month: 'long', day: 'numeric' });
                            }
                        }
                    },
                    legend: {
                        labels: {
                            color: labelColor
                        }
                    },
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
                        grid: {
                            color: gridColor
                        },
                        type: 'time',
                        time: {
                            tooltipFormat: dayFormat,
                            unit: 'day'
                        },
                        display: true,
                        title: {
                            display: true,
                            text: 'Date',
                            color: labelColor
                        },
                        ticks: {
                            color: ticksColor
                        }
                    },
                    y: {
                        grid: {
                            color: gridColor
                        },
                        title: {
                            display: true,
                            text: 'Count',
                            color: titleColor
                        },
                        ticks: {
                            color: ticksColor
                        }
                    }
                },
            }
        }
    }
});
</script>
