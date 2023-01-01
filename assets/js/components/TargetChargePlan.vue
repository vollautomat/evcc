<template>
	<div class="root">
		<h1>Plan</h1>
		<canvas class="plan" ref="plan"></canvas>
		<div class="prices">
			<div class="me-3">
				<shopicon-regular-powersupply
					class="gridIcon"
					size="m"
				></shopicon-regular-powersupply>
			</div>
			<div v-for="price in priceSlots" :key="price.start" class="price">
				{{ fmtPricePerKWh(price.price, currency).replace("/kWh", "") }}
				<div class="box" :style="priceStyle(price.price)"></div>
			</div>
		</div>
		<hr />
		<div class="chargingSlots">
			<div class="me-3">
				<shopicon-regular-lightning
					class="chargingIcon"
					size="m"
				></shopicon-regular-lightning>
			</div>
			<div
				v-for="chargingSlot in chargingSlots"
				:key="chargingSlot.start"
				class="chargingSlot"
			></div>
		</div>
	</div>
</template>

<script>
import "@h2d2/shopicons/es/regular/lightning";
import "@h2d2/shopicons/es/regular/powersupply";
import formatter from "../mixins/formatter";
import {
	Chart,
	LinearScale,
	CategoryScale,
	PointElement,
	LineElement,
	LineController,
	Filler,
	TimeScale,
	BarController,
	BarElement,
} from "chart.js";
import "chartjs-adapter-luxon";

Chart.register(
	LinearScale,
	CategoryScale,
	PointElement,
	LineElement,
	LineController,
	Filler,
	TimeScale,
	BarController,
	BarElement
);

export default {
	name: "TargetChargePlan",
	mixins: [formatter],
	props: {
		priceSlots: Array,
		co2Slots: Array,
		chargingSlots: Array,
		currency: String,
		energyPrice: Number,
		targetTime: String,
	},
	computed: {
		maxPrice() {
			let result = 0;
			this.priceSlots.forEach(({ price }) => {
				result = Math.max(result, price);
			});
			return result;
		},
	},
	mounted() {
		const RED = "rgb(255, 99, 132)";

		const labels = [
			new Date("2023-01-01T01:00:00"),
			new Date("2023-01-01T02:00:00"),
			new Date("2023-01-01T02:30:00"),
			new Date("2023-01-01T03:00:00"),
			new Date("2023-01-01T08:00:00"),
			new Date("2023-01-01T09:00:00"),
			new Date("2023-01-01T10:00:00"),
		];
		const data = {
			labels: labels,
			datasets: [
				{
					label: "Dataset 1",
					data: [10, 30, 48, 20, 25, 44, 10],
					borderColor: RED,
					backgroundColor: RED,
					borderRadius: 10,
					inflateAmount: 20,
					fill: true,
					stepped: true,
				},
			],
		};
		const config = {
			type: "bar",
			data: data,
			options: {
				responsive: true,
				scales: {
					x: {
						type: "time",
						position: "left",
					},
				},
			},
		};
		new Chart(this.$refs.plan, config);
	},
	methods: {
		priceStyle(price) {
			return { height: `${(100 / this.maxPrice) * price}%` };
		},
	},
};
</script>

<style scoped>
.root {
	overflow: hidden;
	--height: 80px;
}
.prices {
	display: flex;
	height: var(--height);
	justify-content: stretch;
	align-items: flex-end;
}
.price {
	flex-basis: 0;
	flex-grow: 1;
	flex-shrink: 1;
	margin: 4px;
	text-align: center;
	height: 100%;
	display: flex;
	justify-content: flex-end;
	flex-direction: column;
}
.box {
	background: linear-gradient(0deg, #999, #000);
	background-size: 100% var(--height);
	background-position: bottom;
	border-radius: 6px 6px 0 0;
}
.chargingSlots {
	display: flex;
	align-items: center;
}
.chargingSlot {
	margin-left: 34%;
	width: 30%;
	background-color: var(--evcc-darker-green);
	height: 6px;
	border-radius: 6px;
}
.chargingIcon {
	color: var(--evcc-darker-green);
}
.gridIcon {
	color: #999;
}
hr {
	border: none;
	border-bottom: 2px solid black;
}
</style>
