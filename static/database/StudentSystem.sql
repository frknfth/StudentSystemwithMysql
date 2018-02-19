-- phpMyAdmin SQL Dump
-- version 4.7.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Feb 16, 2018 at 02:01 PM
-- Server version: 5.6.38
-- PHP Version: 7.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `StudentSystem`
--

-- --------------------------------------------------------

--
-- Table structure for table `Class`
--

CREATE TABLE `Class` (
  `Id` int(5) NOT NULL,
  `Name` varchar(40) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `Class`
--

INSERT INTO `Class` (`Id`, `Name`) VALUES
(4, 'Bil 372'),
(1, 'test3');

-- --------------------------------------------------------

--
-- Table structure for table `Student`
--

CREATE TABLE `Student` (
  `Student_ID` int(11) NOT NULL,
  `Student_Name` varchar(30) NOT NULL,
  `Student_Age` int(3) NOT NULL,
  `Student_UniversityID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `Student`
--

--
-- Triggers `Student`
--
DELIMITER $$
CREATE TRIGGER `after_delete_student` AFTER DELETE ON `Student` FOR EACH ROW UPDATE University
SET University.University_RecordedStudent =University.University_RecordedStudent-1
WHERE University.University_ID = old.Student_UniversityID
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `after_insert_student` AFTER INSERT ON `Student` FOR EACH ROW UPDATE University
SET University.University_RecordedStudent =University.University_RecordedStudent+1
WHERE University.University_ID = new.Student_UniversityID
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `StudentClass`
--

CREATE TABLE `StudentClass` (
  `s_id` int(4) NOT NULL,
  `c_id` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `StudentClass`
--


-- --------------------------------------------------------

--
-- Table structure for table `Teacher`
--

CREATE TABLE `Teacher` (
  `Id` int(11) NOT NULL,
  `Name` varchar(40) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `Teacher`
--


-- --------------------------------------------------------

--
-- Table structure for table `TeacherClass`
--

CREATE TABLE `TeacherClass` (
  `t_id` int(4) NOT NULL,
  `c_id` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `TeacherClass`
--


-- --------------------------------------------------------

--
-- Table structure for table `University`
--

CREATE TABLE `University` (
  `University_ID` int(11) NOT NULL,
  `University_Name` varchar(30) NOT NULL,
  `University_Capacity` int(5) NOT NULL,
  `University_RecordedStudent` int(5) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `University`
--

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Class`
--
ALTER TABLE `Class`
  ADD PRIMARY KEY (`Id`),
  ADD UNIQUE KEY `Name` (`Name`);

--
-- Indexes for table `Student`
--
ALTER TABLE `Student`
  ADD PRIMARY KEY (`Student_ID`),
  ADD KEY `Student_UniversityID` (`Student_UniversityID`);

--
-- Indexes for table `StudentClass`
--
ALTER TABLE `StudentClass`
  ADD KEY `student-class_ibfk_1` (`s_id`),
  ADD KEY `student-class_ibfk_2` (`c_id`);

--
-- Indexes for table `Teacher`
--
ALTER TABLE `Teacher`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `TeacherClass`
--
ALTER TABLE `TeacherClass`
  ADD KEY `teacher-class_ibfk_1` (`t_id`),
  ADD KEY `teacher-class_ibfk_2` (`c_id`);

--
-- Indexes for table `University`
--
ALTER TABLE `University`
  ADD PRIMARY KEY (`University_ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Class`
--
ALTER TABLE `Class`
  MODIFY `Id` int(5) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `Student`
--
ALTER TABLE `Student`
  MODIFY `Student_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `Teacher`
--
ALTER TABLE `Teacher`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `University`
--
ALTER TABLE `University`
  MODIFY `University_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Student`
--
ALTER TABLE `Student`
  ADD CONSTRAINT `student_ibfk_1` FOREIGN KEY (`Student_UniversityID`) REFERENCES `University` (`University_ID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `StudentClass`
--
ALTER TABLE `StudentClass`
  ADD CONSTRAINT `student-class_ibfk_1` FOREIGN KEY (`s_id`) REFERENCES `Student` (`Student_ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `student-class_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `Class` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `TeacherClass`
--
ALTER TABLE `TeacherClass`
  ADD CONSTRAINT `teacher-class_ibfk_1` FOREIGN KEY (`t_id`) REFERENCES `Teacher` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `teacher-class_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `Class` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
