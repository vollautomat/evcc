<template>
	<form @submit.prevent="setTargetTime">
		<div class="mt-4">
			<div class="form-group d-lg-flex align-items-baseline mb-2 justify-content-between">
				<div v-if="targetTime" class="container px-0">
					<div class="row my-3" :class="{ 'opacity-50': state === 'disabled' }">
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="Mo." />
						</div>
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="10:45" />
						</div>
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="50%" />
						</div>
						<div class="col d-flex align-items-center justify-content-between">
							<button
								v-if="state === 'once'"
								type="button"
								class="btn btn-sm text-primary"
								@click.prevent="state = 'repeat'"
							>
								<shopicon-regular-checkbox></shopicon-regular-checkbox>
							</button>
							<button
								v-if="state === 'repeat'"
								type="button"
								class="btn btn-sm text-primary"
								@click.prevent="state = 'disabled'"
							>
								<shopicon-regular-infinite></shopicon-regular-infinite>
							</button>
							<button
								v-if="state === 'disabled'"
								type="button"
								class="btn btn-sm text-muted"
								@click.prevent="state = 'once'"
							>
								<shopicon-regular-square></shopicon-regular-square>
							</button>
							<button type="button" class="btn btn-sm btn-link text-muted">
								<shopicon-regular-trash></shopicon-regular-trash>
							</button>
						</div>
					</div>
					<div class="row my-3">
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="Sa." />
						</div>
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="8:00" />
						</div>
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="75%" />
						</div>
						<div class="col d-flex align-items-center justify-content-between">
							<button type="button" class="btn btn-sm text-primary">
								<shopicon-regular-checkbox></shopicon-regular-checkbox>
							</button>
							<button type="button" class="btn btn-sm btn-link text-muted">
								<shopicon-regular-trash></shopicon-regular-trash>
							</button>
						</div>
					</div>
					<div class="row my-3">
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="Di.-Do." />
						</div>
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="6:00" />
						</div>
						<div class="col-3">
							<input type="text" class="form-control mx-0" value="55%" />
						</div>
						<div class="col d-flex align-items-center justify-content-between">
							<button type="button" class="btn btn-sm text-primary">
								<shopicon-regular-infinite></shopicon-regular-infinite>
							</button>
							<button type="button" class="btn btn-sm btn-link text-muted">
								<shopicon-regular-trash></shopicon-regular-trash>
							</button>
						</div>
					</div>
					<div class="d-flex justify-content-end">
						<button class="btn btn-sm d-flex align-items-center">
							<shopicon-regular-plus size="s"></shopicon-regular-plus>
						</button>
					</div>
				</div>
				<div v-else>
					<p>
						No charging target set. Set your departure and charge goals; evcc computes
						the optimal charging schedule.
					</p>
					<button class="btn btn-outline-primary">Set charging target</button>
				</div>
				<!--
				<div class="d-flex justify-content-between date-selection">
					<select
						:id="`targetTimeLabel${id}`"
						v-model="selectedDay"
						class="form-select me-2"
						data-testid="target-day"
					>
						<option v-for="opt in dayOptions()" :key="opt.value" :value="opt.value">
							{{ opt.name }}
						</option>
					</select>
					<input
						v-model="selectedTime"
						type="time"
						class="form-control ms-2 time-selection"
						:step="60 * 5"
						required
						data-testid="target-time"
					/>
				</div>
				-->
			</div>
			<div v-if="targetTime">
				<hr />
				<h5>PREVIEW</h5>
				<!--
			<p class="mb-0">
				<span v-if="timeInThePast" class="d-block text-danger mb-1">
					{{ $t("main.targetCharge.targetIsInThePast") }}
				</span>
				<span v-if="!socBasedCharging && !targetEnergy" class="d-block text-danger mb-1">
					{{ $t("main.targetCharge.targetEnergyRequired") }}
				</span>
				<span v-if="timeTooFarInTheFuture" class="d-block text-secondary mb-1">
					{{ $t("main.targetCharge.targetIsTooFarInTheFuture") }}
				</span>
				<span v-if="costLimitExists" class="d-block text-secondary mb-1">
					{{
						$t("main.targetCharge.costLimitIgnore", {
							limit: costLimitText,
						})
					}}
				</span>
				<span v-if="['off', 'now'].includes(mode)" class="d-block text-secondary mb-1">
					{{ $t("main.targetCharge.onlyInPvMode") }}
				</span>
				&nbsp;
			</p>
			-->
				<TargetChargePlan v-if="targetChargePlanProps" v-bind="targetChargePlanProps" />
				<!--
				<div class="d-flex justify-content-between mt-3">
					<button
						type="button"
						class="btn btn-outline-secondary"
						:disabled="!targetTime"
						@click="removeTargetTime"
					>
						{{ $t("main.targetCharge.remove") }}
					</button>
					<button type="submit" class="btn btn-primary" :disabled="timeInThePast">
						<span v-if="targetTime">
							{{ $t("main.targetCharge.update") }}
						</span>
						<span v-else>
							{{ $t("main.targetCharge.activate") }}
						</span>
					</button>
				</div>-->
			</div>
		</div>
	</form>
</template>

<script>
import "@h2d2/shopicons/es/regular/square";
import "@h2d2/shopicons/es/regular/trash";
import "@h2d2/shopicons/es/regular/infinite";
import "@h2d2/shopicons/es/regular/checkbox";
import "@h2d2/shopicons/es/regular/plus";
import { CO2_TYPE } from "../units";
import TargetChargePlan from "./TargetChargePlan.vue";
import api from "../api";

import formatter from "../mixins/formatter";

const DEFAULT_TARGET_TIME = "7:00";
const LAST_TARGET_TIME_KEY = "last_target_time";

export default {
	name: "TargetCharge",
	components: { TargetChargePlan },
	mixins: [formatter],
	props: {
		id: [String, Number],
		planActive: Boolean,
		targetTime: String,
		targetSoc: Number,
		targetEnergy: Number,
		socBasedCharging: Boolean,
		disabled: Boolean,
		smartCostLimit: Number,
		smartCostType: String,
		currency: String,
		mode: String,
	},
	emits: ["target-time-updated", "target-time-removed"],
	data: function () {
		return {
			selectedDay: null,
			selectedTime: null,
			plan: {},
			tariff: {},
			activeTab: "time",
			loading: false,
			state: "disabled",
		};
	},
	computed: {
		timeTabActive: function () {
			return this.activeTab === "time";
		},
		priceTabActive: function () {
			return this.activeTab === "price";
		},
		targetChargeEnabled: function () {
			return this.targetTime;
		},
		timeInThePast: function () {
			const now = new Date();
			return now >= this.selectedTargetTime;
		},
		timeTooFarInTheFuture: function () {
			if (this.tariff?.rates) {
				const lastRate = this.tariff.rates[this.tariff.rates.length - 1];
				if (lastRate?.end) {
					const end = new Date(lastRate.end);
					return this.selectedTargetTime >= end;
				}
			}
			return false;
		},
		selectedTargetTime: function () {
			return new Date(`${this.selectedDay}T${this.selectedTime || "00:00"}`);
		},
		targetEnergyFormatted: function () {
			return this.fmtKWh(this.targetEnergy * 1e3, true, true, 1);
		},
		targetChargePlanProps: function () {
			const targetTime = this.selectedTargetTime;
			const { rates } = this.tariff;
			const { duration, plan } = this.plan;
			const { currency, smartCostType } = this;
			return rates ? { duration, rates, plan, targetTime, currency, smartCostType } : null;
		},
		tariffLowest: function () {
			return this.tariff?.rates.reduce((res, slot) => {
				return Math.min(res, slot.price);
			}, Number.MAX_VALUE);
		},
		tariffHighest: function () {
			return this.tariff?.rates.reduce((res, slot) => {
				return Math.max(res, slot.price);
			}, 0);
		},
		costLimitExists: function () {
			return this.smartCostLimit !== 0;
		},
		costLimitText: function () {
			if (this.isCo2) {
				return this.$t("main.targetCharge.co2Limit", {
					co2: this.fmtCo2Short(this.smartCostLimit),
				});
			}
			return this.$t("main.targetCharge.priceLimit", {
				price: this.fmtPricePerKWh(this.smartCostLimit, this.currency, true),
			});
		},
		isCo2() {
			return this.smartCostType === CO2_TYPE;
		},
	},
	watch: {
		targetTime() {
			this.initInputFields();
			this.updatePlan();
		},
		selectedTargetTime() {
			this.updatePlan();
		},
		targetSoc() {
			this.updatePlan();
		},
		targetEnergy() {
			this.updatePlan();
		},
	},
	mounted() {
		this.initInputFields();
		this.updatePlan();
	},
	methods: {
		updatePlan: async function () {
			if (
				!this.timeInThePast &&
				(this.targetEnergy || this.targetSoc) &&
				!isNaN(this.selectedTargetTime) &&
				!this.loading
			) {
				try {
					this.loading = true;
					this.plan = (
						await api.get(`loadpoints/${this.id}/target/plan`, {
							params: { targetTime: this.selectedTargetTime },
						})
					).data.result;

					const tariffRes = await api.get(`tariff/planner`, {
						validateStatus: function (status) {
							return status >= 200 && status < 500;
						},
					});
					this.tariff = tariffRes.status === 404 ? { rates: [] } : tariffRes.data.result;
				} catch (e) {
					console.error(e);
				} finally {
					this.loading = false;
				}
			}
		},
		defaultDate: function () {
			const [hours, minutes] = (
				window.localStorage[LAST_TARGET_TIME_KEY] || DEFAULT_TARGET_TIME
			).split(":");

			const target = new Date();
			target.setSeconds(0);
			target.setMinutes(minutes);
			target.setHours(hours);
			// today or tomorrow?
			const isInPast = target < new Date();
			if (isInPast) {
				target.setDate(target.getDate() + 1);
			}
			return target;
		},
		initInputFields: function () {
			let date = this.defaultDate();
			let targetTimeInTheFuture = new Date(this.targetTime) > new Date();
			if (this.targetChargeEnabled && targetTimeInTheFuture) {
				date = new Date(this.targetTime);
			}
			this.selectedDay = this.fmtDayString(date);
			this.selectedTime = this.fmtTimeString(date);
		},
		dayOptions: function () {
			const options = [];
			const date = new Date();
			const labels = [
				this.$t("main.targetCharge.today"),
				this.$t("main.targetCharge.tomorrow"),
			];
			for (let i = 0; i < 7; i++) {
				const dayNumber = date.toLocaleDateString(this.$i18n.locale, {
					month: "short",
					day: "numeric",
				});
				const dayName =
					labels[i] || date.toLocaleDateString(this.$i18n.locale, { weekday: "long" });
				options.push({
					value: this.fmtDayString(date),
					name: `${dayNumber} (${dayName})`,
				});
				date.setDate(date.getDate() + 1);
			}
			return options;
		},
		setTargetTime: function () {
			try {
				const hours = this.selectedTargetTime.getHours();
				const minutes = this.selectedTargetTime.getMinutes();
				window.localStorage[LAST_TARGET_TIME_KEY] = `${hours}:${minutes}`;
			} catch (e) {
				console.warn(e);
			}
			this.$emit("target-time-updated", this.selectedTargetTime);
		},
		removeTargetTime: function () {
			this.$emit("target-time-removed");
		},
	},
};
</script>

<style scoped>
@media (min-width: 992px) {
	.date-selection {
		width: 370px;
	}
}
.time-selection {
	flex-basis: 200px;
}
h5 {
	position: relative;
	display: inline-block;
	background-color: white;
	top: -25px;
	left: calc(50% - 50px);
	padding: 0 0.5rem;
	font-weight: normal;
	color: var(--bs-gray);
	margin-bottom: -4rem;
}
</style>
