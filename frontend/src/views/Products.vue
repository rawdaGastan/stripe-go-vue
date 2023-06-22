<template>
	<v-container class="pa-1" fluid>
		<v-row>
			<v-col cols="12" sm="8">
				<h1 style="text-align:center; font-size:2.5em;">Products</h1>
				<v-row no-gutters>
					<v-col class="item" v-for="item in products" cols="12" sm="6" md="3" v-bind:key="item">
						<div class="section">
							<div class="product">
								<img src="https://i.imgur.com/EHyR2nP.png" alt="The cover of Stubborn Attachments" />
								<div class="description">
									<h3>{{ item.name }}</h3>
									<h5>${{ item.price }}</h5>
								</div>
							</div>
							<v-form @submit.prevent="addProduct(item)">
								<button type="submit">Add</button>
							</v-form>
						</div>
					</v-col>
				</v-row>
			</v-col>
			<v-divider vertical />
			<v-col cols="12" sm="4" class="d-flex justify-center align-center">
				<h1 v-if="cart.length == 0" style="text-align:center; font-size:1.5em;">Your cart is empty, please add some items
				</h1>
				<div v-else class="section" style="height: 500px; width:80%">
					<v-list>
						<v-list-subheader>Cart</v-list-subheader>

						<v-list-item v-for="(item, i) in cart" :key="i">
							<div class="product">
								<img src="https://i.imgur.com/EHyR2nP.png" />
								<div class="description">
									<h3>{{ item.amount }}× {{ item.name }}</h3>
									<h5>${{ item.price }}</h5>
								</div>
							</div>
						</v-list-item>
					</v-list>


					<v-form @submit.prevent="checkout">
						<button type="submit">Checkout ({{ total }}$)</button>
					</v-form>
				</div>
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
import stripeService from "@/services/stripeService";
import { ref, onMounted } from "vue";

export default {
	name: "Products",
	setup() {
		const products = ref([]);
		const cart = ref([]);
		const outCart = ref([]);
		const total = ref(0);

		const getProducts = () => {
			stripeService
				.getProducts()
				.then((response) => {
					products.value = response.data.data;
				})
				.catch(() => {
					products.value = [];
				});
		};

		const addProduct = (item) => {
			outCart.value.push({ "product_id": item.id, "amount": 1 });
			total.value += item.price;

			let added = false;
			cart.value.forEach((i) => {
				if (i.name == item.name) {
					i.amount += 1;
					i.price += item.price;
					added = true;
				}
			});

			if (!added) {
				cart.value.push({ "name": item.name, "amount": 1, "price": item.price });
			}
		};

		const checkout = () => {
			let success_url = "http://localhost:8080/success";
			let failure_url = "http://localhost:8080/canceled";
			stripeService.checkout(outCart.value, success_url, failure_url).then((response) => {
				window.location.href = response.data.data;
			});
		};

		onMounted(() => {
			getProducts();
		});
		return { getProducts, checkout, addProduct, products, cart, outCart, total };
	},
};
</script>

<style scoped>
.section {
	background: #ffffff;
	display: flex;
	flex-direction: column;
	border-radius: 6px;
	justify-content: space-between;
	margin: 10%;
}

p {
	font-style: normal;
	font-weight: 500;
	font-size: 14px;
	line-height: 20px;
	letter-spacing: -0.154px;
	color: #242d60;
	height: 100%;
	width: 100%;
	padding: 0 20px;
	display: flex;
	align-items: center;
	justify-content: center;
	box-sizing: border-box;
	padding: 10%;
}

.product {
	display: flex;
}

.description {
	display: flex;
	flex-direction: column;
	justify-content: center;
}

h3,
h5 {
	font-style: normal;
	font-weight: 500;
	font-size: 14px;
	line-height: 20px;
	letter-spacing: -0.154px;
	color: #242d60;
	margin: 0;
}

h5 {
	opacity: 0.5;
}

img {
	border-radius: 6px;
	margin: 10px;
	width: 54px;
	height: 57px;
}


button {
	height: 36px;
	background: #556cd6;
	color: white;
	width: 100%;
	font-size: 14px;
	border: 0;
	font-weight: 500;
	cursor: pointer;
	letter-spacing: 0.6;
	border-radius: 0 0 6px 6px;
	transition: all 0.2s ease;
	box-shadow: 0px 4px 5.5px 0px rgba(0, 0, 0, 0.07);
}

button:hover {
	opacity: 0.8;
}
</style>
