<template>
	<div class="mb-2 entry" :class="{ 'evcc-gray': !active }">
		<div class="d-flex justify-content-between">
			<span class="d-flex flex-nowrap">
				<BatteryIcon v-if="isBattery" v-bind="iconProps" />
				<VehicleIcon v-else-if="isVehicle" v-bind="iconProps" />
				<component :is="`shopicon-regular-${icon}`" v-else></component>
			</span>
			<div class="d-block text-nowrap flex-grow-1 ms-3 text-truncate">
				{{ name }}
			</div>
			<span class="text-end text-nowrap ps-1 fw-bold d-flex">
				<div
					ref="details"
					class="fw-normal"
					:class="{ 'text-decoration-underline': detailsClickable }"
					data-testid="energyflow-entry-details"
					data-bs-toggle="tooltip"
					:tabindex="detailsClickable ? 0 : undefined"
					@click="detailsClicked"
				>
					<AnimatedNumber v-if="!isNaN(details)" :to="details" :format="detailsFmt" />
				</div>
				<div ref="power" class="power" data-bs-toggle="tooltip" @click="powerClicked">
					<AnimatedNumber ref="powerNumber" :to="power" :format="kw" />
				</div>
			</span>
		</div>
		<div v-if="name == $t('main.energyflow.gridImport')" class="ms-4 ps-3">
			<div class="d-flex justify-content-between">
				<span>Status</span>
				<span class="text-end text-nowrap ps-1 d-flex">{{consLimitStatus}}</span>
				<span class="text-end text-nowrap ps-1 fw-bold d-flex">
					<label>{{consLimitDuration}} s</label>
					<div ref="climit" class="power">
						<AnimatedNumber ref="climitNumber" :to="consLimitValue" :format="kw" />
					</div>
				</span>
			</div>
		</div>
		<div v-if="$slots.subline" class="ms-4 ps-3">
			<slot name="subline" />
		</div>
	</div>
</template>

<script>
import "@h2d2/shopicons/es/regular/powersupply";
import "@h2d2/shopicons/es/regular/sun";
import "@h2d2/shopicons/es/regular/home";
import Tooltip from "bootstrap/js/dist/tooltip";
import BatteryIcon from "./BatteryIcon.vue";
import formatter from "../../mixins/formatter";
import AnimatedNumber from "../AnimatedNumber.vue";
import VehicleIcon from "../VehicleIcon";

export default {
	name: "EnergyflowEntry",
	components: { BatteryIcon, AnimatedNumber, VehicleIcon },
	mixins: [formatter],
	props: {
		name: { type: String },
		icon: { type: String },
		iconProps: { type: Object, default: () => ({}) },
		power: { type: Number },
		powerTooltip: { type: Array },
		powerUnit: { type: String },
		details: { type: Number },
		detailsFmt: { type: Function },
		detailsTooltip: { type: Array },
		detailsClickable: { type: Boolean },
		consLimit: { type: Object },
	},
	emits: ["details-clicked"],
	data() {
		return {
			powerTooltipInstance: null,
			detailsTooltipInstance: null,
			nowTrigger: 0,
		};
	},
	computed: {
		active: function () {
			return this.power > 10;
		},
		isBattery: function () {
			return this.icon === "battery";
		},
		isVehicle: function () {
			return this.icon === "vehicle";
		},
		consLimitStatus: function () {
			if ( ! this.consLimit )
				return "Unknown";

			switch (this.consLimit.status) {
				case 0: return "Unlimited";
				case 1: return "Limited";
				case 2: return "Failsafe";
				default: return "Unknown";
			}
		},
		consLimitValue: function () {
			if ( ! this.consLimit || this.consLimit.status == 0)
				return 0;

			if ( this.consLimit.status == 1 )
				return this.consLimit.consumptionLimit.Value;
			else if ( this.consLimit.status == 2 )
				return this.consLimit.failsafeLimit;
			else
				return 0;
		},
		consLimitDuration: function () {
			if ( ! this.consLimit || this.consLimit.status == 0 )
				return 0;

			this.nowTrigger++;

			var duration = 0;
			if ( this.consLimit.status == 1 ) {
				duration = new Date( this.consLimit.statusUpdated ).getTime()
						 + this.consLimit.consumptionLimit.Duration / 1000000
						 - Date.now();
			}
			else if ( this.consLimit.status == 2 ) {
				duration = new Date( this.consLimit.statusUpdated ).getTime()
						 + this.consLimit.failsafeDuration / 1000000
						 - Date.now();
			} 
			this.duration = Math.round( duration / 1000 + 0.5, 0 );
			return Math.round( duration / 1000 + 0.5, 0 );
		}
	},
	watch: {
		powerTooltip(newVal, oldVal) {
			if (JSON.stringify(newVal) !== JSON.stringify(oldVal)) {
				this.updatePowerTooltip();
			}
		},
		detailsTooltip(newVal, oldVal) {
			if (JSON.stringify(newVal) !== JSON.stringify(oldVal)) {
				this.updateDetailsTooltip();
			}
		},
		powerInKw(newVal, oldVal) {
			// force update if unit changes but not the value
			if (newVal !== oldVal) {
				this.$refs.powerNumber.forceUpdate();
			}
		},
	},
	mounted: function () {
		this.updatePowerTooltip();
		this.updateDetailsTooltip();

		// setInterval(function() {
		// 	if ( this.consLimit ) {
		// 		console.log( Date.now() );

		// 		var duration = new Date( this.consLimit.statusUpdated ).getTime()
		// 					+ this.consLimit.consumptionLimit.Duration / 1000000
		// 					- Date.now();
		// 		this.duration = Math.round( duration / 1000 + 0.5, 0 );
		// 	}
		// }, 1000);
	},
	methods: {
		kw: function (watt) {
			return this.fmtW(watt, this.powerUnit);
		},
		updatePowerTooltip() {
			this.powerTooltipInstance = this.updateTooltip(
				this.powerTooltipInstance,
				this.powerTooltip,
				this.$refs.power
			);
		},
		updateDetailsTooltip() {
			if (this.detailsClickable) {
				return;
			}
			this.detailsTooltipInstance = this.updateTooltip(
				this.detailsTooltipInstance,
				this.detailsTooltip,
				this.$refs.details
			);
		},
		updateTooltip: function (instance, content, ref) {
			if (!Array.isArray(content) || !content.length) {
				if (instance) {
					instance.dispose();
				}
				return;
			}
			let newInstance = instance;
			if (!newInstance) {
				newInstance = new Tooltip(ref, { html: true, title: " " });
			}
			const html = `<div class="text-end">${content.join("<br/>")}</div>`;
			newInstance.setContent({ ".tooltip-inner": html });
			return newInstance;
		},
		powerClicked: function ($event) {
			if (this.powerTooltip) {
				$event.stopPropagation();
			}
		},
		detailsClicked: function ($event) {
			if (this.detailsClickable || this.detailsTooltip) {
				$event.stopPropagation();
			}
			if (this.detailsClickable) {
				this.$emit("details-clicked");
			}
		},
	},
};
</script>
<style scoped>
.entry {
	transition: color var(--evcc-transition-medium) linear;
}
.power {
	min-width: 75px;
}
</style>
