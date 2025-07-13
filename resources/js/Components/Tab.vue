<template>
    <div class="border-b border-gray-200 dark:border-gray-700 mt-20">
        <ul
            class="flex flex-wrap -mb-px text-sm font-medium text-center text-gray-500 dark:text-gray-400"
        >
            <li class="me-2">
                <Link href="/" v-bind:class="itemClass(page.url === '/')">
                    <Bars3BottomLeftIcon class="size-6" />
                    &nbsp; Raw Data
                </Link>
            </li>
            <li class="me-2">
                <Link
                    href="/url"
                    v-bind:class="itemClass(page.url.startsWith('/url'))"
                >
                    <LinkIcon class="size-6" />
                    &nbsp; URL
                </Link>
            </li>
            <li class="me-2">
                <Link
                    href="/phone-number"
                    v-bind:class="
                        itemClass(page.url.startsWith('/phone-number'))
                    "
                >
                    <PhoneIcon class="size-6" />
                    &nbsp; Phone Number
                </Link>
            </li>
            <li class="me-2">
                <a
                    href="#"
                    v-bind:class="itemClass(page.url.startsWith('/v-card'))"
                >
                    <IdentificationIcon class="size-6" />
                    &nbsp; vCard
                </a>
            </li>
            <li class="me-2">
                <a
                    href="#"
                    v-bind:class="itemClass(page.url.startsWith('/email'))"
                >
                    <EnvelopeIcon class="size-6" />
                    &nbsp; Email
                </a>
            </li>
        </ul>
    </div>
</template>

<script setup lang="ts">
import {
    Bars3BottomLeftIcon,
    EnvelopeIcon,
    IdentificationIcon,
    LinkIcon,
    PhoneIcon,
} from "@heroicons/vue/24/outline";
import { Link, usePage } from "@inertiajs/vue3";

const page = usePage();

const itemClass = (active: boolean) => {
    if (active) {
        return "inline-flex items-center justify-center p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-blue-500 dark:border-blue-500 group";
    }

    return "inline-flex items-center justify-center p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300 group";
};
</script>
