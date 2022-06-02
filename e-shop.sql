-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 02, 2022 at 06:10 PM
-- Server version: 10.6.7-MariaDB
-- PHP Version: 8.1.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `e-shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `cart_details`
--

CREATE TABLE `cart_details` (
  `id` int(11) NOT NULL,
  `cart_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `description`) VALUES
(1, 'Barang Langka', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `merchants`
--

CREATE TABLE `merchants` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `address` text NOT NULL,
  `phone_number` varchar(20) NOT NULL,
  `tagline` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `merchants`
--

INSERT INTO `merchants` (`id`, `name`, `address`, `phone_number`, `tagline`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Toko 1', 'Jalan Toko 1 Di Bandung ', '01821387183', 'Bersama kami membangun negeri', 'Deskripsi toko blablabla', '2022-06-01 14:29:26', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `merchants_id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` int(11) NOT NULL,
  `discounted_price` int(11) DEFAULT NULL,
  `description` text NOT NULL,
  `stocks` int(11) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `merchants_id`, `category_id`, `name`, `price`, `discounted_price`, `description`, `stocks`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'Produk toko 1', 100000, NULL, 'Deskripsi', 100, '2022-06-01 14:29:47', NULL),
(2, 1, 1, 'Keripik Kaca', 200000, NULL, 'Hello World', 200, '2022-06-02 17:23:45', NULL),
(3, 1, 1, 'Keripik Kaca 2', 200000, NULL, 'Hello World', 200, '2022-06-02 17:23:45', NULL),
(4, 1, 1, 'Keripik Kaca 3', 20000, NULL, 'Hello World', 200, '2022-06-02 17:23:45', NULL),
(5, 1, 1, 'Keripik Kaca 4', 200000, NULL, 'Hello World', 200, '2022-06-02 17:23:45', NULL),
(6, 1, 1, 'Keripik Kaca 5', 1200000, NULL, 'Hello World', 200, '2022-06-02 17:23:45', NULL),
(7, 1, 1, 'Keripik Kaca 6', 20012000, NULL, 'Hello World', 200, '2022-06-02 17:23:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `product_images`
--

CREATE TABLE `product_images` (
  `id` int(11) NOT NULL,
  `image_url` varchar(50) NOT NULL,
  `product_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_images`
--

INSERT INTO `product_images` (`id`, `image_url`, `product_id`) VALUES
(1, '9_1638280692144_ic_profile_default.png', 1);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `payment_url` varchar(255) NOT NULL,
  `status` varchar(50) NOT NULL DEFAULT 'pending' COMMENT 'pending, paid, cancelled, failed',
  `deadline` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `total_price`, `payment_url`, `status`, `deadline`, `created_at`, `updated_at`) VALUES
(1, 24, 100000, 'https://payment-ni-bos', 'pending', '2022-06-02 14:07:21', '2022-06-02 21:07:39', NULL),
(2, 24, 0, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:01:22', '2022-06-02 22:01:22'),
(3, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:02:39', '2022-06-02 22:02:39'),
(4, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:03:59', '2022-06-02 22:03:59'),
(5, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:06:41', '2022-06-02 22:06:41'),
(6, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:07:54', '2022-06-02 22:07:54'),
(7, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:10:45', '2022-06-02 22:10:45'),
(8, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:11:48', '2022-06-02 22:11:48'),
(9, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:12:30', '2022-06-02 22:12:30'),
(10, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:13', '2022-06-02 22:13:13'),
(11, 24, 6000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:37', '2022-06-02 22:13:37'),
(12, 24, 4000000, '', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:26:39', '2022-06-02 22:26:39'),
(13, 24, 100000, 'https://payment-ni-bos', 'pending', '2022-06-02 14:07:21', '2022-06-02 21:07:39', NULL),
(14, 24, 0, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:01:22', '2022-06-02 22:01:22'),
(15, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:02:39', '2022-06-02 22:02:39'),
(16, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:03:59', '2022-06-02 22:03:59'),
(17, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:06:41', '2022-06-02 22:06:41'),
(18, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:07:54', '2022-06-02 22:07:54'),
(19, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:10:45', '2022-06-02 22:10:45'),
(20, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:11:48', '2022-06-02 22:11:48'),
(21, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:12:30', '2022-06-02 22:12:30'),
(22, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:13', '2022-06-02 22:13:13'),
(23, 24, 6000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:37', '2022-06-02 22:13:37'),
(24, 24, 4000000, '', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:26:39', '2022-06-02 22:26:39'),
(25, 24, 100000, 'https://payment-ni-bos', 'pending', '2022-06-02 14:07:21', '2022-06-02 21:07:39', NULL),
(26, 24, 0, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:01:22', '2022-06-02 22:01:22'),
(27, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:02:39', '2022-06-02 22:02:39'),
(28, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:03:59', '2022-06-02 22:03:59'),
(29, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:06:41', '2022-06-02 22:06:41'),
(30, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:07:54', '2022-06-02 22:07:54'),
(31, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:10:45', '2022-06-02 22:10:45'),
(32, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:11:48', '2022-06-02 22:11:48'),
(33, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:12:30', '2022-06-02 22:12:30'),
(34, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:13', '2022-06-02 22:13:13'),
(35, 24, 6000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:37', '2022-06-02 22:13:37'),
(36, 24, 4000000, '', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:26:39', '2022-06-02 22:26:39'),
(37, 24, 100000, 'https://payment-ni-bos', 'pending', '2022-06-02 14:07:21', '2022-06-02 21:07:39', NULL),
(38, 24, 0, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:01:22', '2022-06-02 22:01:22'),
(39, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:02:39', '2022-06-02 22:02:39'),
(40, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:03:59', '2022-06-02 22:03:59'),
(41, 24, 100000, 'https://hello.world', '', '0001-01-02 07:07:12', '2022-06-02 22:06:41', '2022-06-02 22:06:41'),
(42, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:07:54', '2022-06-02 22:07:54'),
(43, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:10:45', '2022-06-02 22:10:45'),
(44, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:11:48', '2022-06-02 22:11:48'),
(45, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:12:30', '2022-06-02 22:12:30'),
(46, 24, 2000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:13', '2022-06-02 22:13:13'),
(47, 24, 6000000, 'https://hello.world', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:13:37', '2022-06-02 22:13:37'),
(48, 24, 4000000, '', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:26:39', '2022-06-02 22:26:39'),
(49, 24, 4000000, '', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:35:36', '2022-06-02 22:35:36'),
(50, 24, 4000000, 'https://app.sandbox.midtrans.com/snap/v2/vtweb/1f4700d0-1ffe-4227-ad39-c39942897699', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:36:04', '2022-06-02 22:36:04'),
(51, 24, 4000000, 'https://app.sandbox.midtrans.com/snap/v2/vtweb/cec50d9d-c7ab-4b8c-a48b-eb5425d911e8', 'pending', '0001-01-02 07:07:12', '2022-06-02 22:37:41', '2022-06-02 22:37:42'),
(52, 24, 4000000, 'https://app.sandbox.midtrans.com/snap/v2/vtweb/e28f4b53-9e53-44bc-96d3-221376a24d20', 'pending', '2022-06-03 22:38:36', '2022-06-02 22:38:36', '2022-06-02 22:38:36');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `id` int(11) NOT NULL,
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`id`, `transaction_id`, `product_id`, `quantity`) VALUES
(1, 1, 3, 100),
(3, 1, 5, 50),
(5, 4, 1, 10000),
(6, 5, 1, 20),
(7, 6, 1, 20),
(8, 7, 1, 20),
(9, 8, 1, 20),
(10, 9, 1, 20),
(11, 10, 1, 20),
(12, 11, 1, 20),
(13, 11, 2, 20),
(14, 50, 2, 20),
(15, 51, 2, 20),
(16, 52, 2, 20);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `address` text NOT NULL,
  `phone_number` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `address`, `phone_number`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Jonhon', 'Kumala', '0192019301', 'jhonny@gmail.com', '$2a$04$eb1YVfFq8A137rO6tsijAunS5tyiIP/IjCgSTtBLYCRPDpSMGx4j6', '2022-05-31 13:39:28', '2022-05-31 13:39:28'),
(4, 'Jonhon', 'Kumala', '0192011419301', 'jhonny123@gmail.com', '$2a$04$.JvrZjfzkZ3ovgQlTs1qQO.s17ADsV2.1GG.427CGx7CS9Grgxnle', '2022-05-31 13:50:15', '2022-05-31 13:50:15'),
(5, 'Ini Upl 13131', 'Jalan', 'Hqkqssqsqq', 'sqoskq@.sqmssqqsqsq', '$2a$04$DgSSb5I6vGdXX5v5aCn9z.xjCpanQob56raxGpHV9gTjxM95/d10O', '2022-06-01 16:00:00', '2022-06-01 16:22:17'),
(18, 'Ini Upla', 'Jalan', '', '', '$2a$04$21HJttLbLfJKYCtPCymZMuFM00Qg6AysuTWpgvyJTJsShocvHGYAS', '2022-06-01 16:18:29', '2022-06-01 16:18:29'),
(20, 'Ini Upla', 'Jalan', 'Hqk', 'sqoskq@.sqmsq', '$2a$04$DxAVtpdKNs6wdoTdY8c1P.c70vxxjXVlU754lMfF9gNf5hDzVv0jy', '2022-06-01 16:18:52', '2022-06-01 16:18:52'),
(22, 'Ini Upl 2', 'Jalan', 'Hqkqsq', 'sqoskq@.sqmsqsqsq', '$2a$04$G8uu2vzlVESvEVoB/XFrle1nzXPWn.KXqZAhmVZpNpO6ZcclO2xii', '2022-06-01 16:19:24', '2022-06-01 16:19:24'),
(23, 'string', 'string', 'string', 'string@string.com', '$2a$04$tz4zY11dNiUKNMotvEx4EOxaj.o.jXshUk.PyCfgfA3jBNZ10XdBe', '2022-06-02 14:44:41', '2022-06-02 14:44:41'),
(24, 'Naufal', 'Jalan', '01281081', 'fawaznaufal23@gmail.com', '$2a$04$xQsYpjBOmeizmfaVXhUv7ODhAho9eDi4JDeWBrMJFd1qiwX.jgIIi', '2022-06-02 14:50:46', '2022-06-02 14:50:46');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `carts_ibfk_1` (`user_id`);

--
-- Indexes for table `cart_details`
--
ALTER TABLE `cart_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `products_id` (`product_id`),
  ADD KEY `cart_details_ibfk_1` (`cart_id`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `merchants`
--
ALTER TABLE `merchants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unique_phone_number` (`phone_number`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `merchants_id` (`merchants_id`),
  ADD KEY `category_id` (`category_id`);

--
-- Indexes for table `product_images`
--
ALTER TABLE `product_images`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_images_ibfk_1` (`product_id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `users_id` (`user_id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `transactions_id` (`transaction_id`),
  ADD KEY `products_id` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone_number` (`phone_number`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `cart_details`
--
ALTER TABLE `cart_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `merchants`
--
ALTER TABLE `merchants`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `product_images`
--
ALTER TABLE `product_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=53;

--
-- AUTO_INCREMENT for table `transaction_details`
--
ALTER TABLE `transaction_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `carts`
--
ALTER TABLE `carts`
  ADD CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `cart_details`
--
ALTER TABLE `cart_details`
  ADD CONSTRAINT `cart_details_ibfk_1` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `cart_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`merchants_id`) REFERENCES `merchants` (`id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `product_images`
--
ALTER TABLE `product_images`
  ADD CONSTRAINT `product_images_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
