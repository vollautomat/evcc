<template>
	<div class="root">
		<h1>Plan</h1>
		<div class="prices">
			<div v-for="price in priceSlots" :key="price.start" class="price">
				{{ fmtPricePerKWh(price.price, currency).replace("/kWh", "") }}
				<div class="box" :style="priceStyle(price.price)"></div>
			</div>
		</div>
	</div>
</template>

<script>
import formatter from "../mixins/formatter";

export default {
	name: "TargetChargePlan",
	mixins: [formatter],
	props: {
		priceSlots: Array,
		co2Slots: Array,
		plannedSlots: Array,
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
}
.prices {
	display: flex;
	height: 100px;
	justify-content: stretch;
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
	background-color: orange;
}
</style>
