

CREATE TABLE `Class` (
  `Id` int(5) NOT NULL,
  `Department` text NOT NULL,
  `Code` int(11) NOT NULL,
  `Name` varchar(40) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `Department` (
  `Id` int(11) NOT NULL,
  `Name` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `Student` (
  `Student_ID` int(11) NOT NULL,
  `Student_Name` varchar(30) NOT NULL,
  `Student_Age` int(3) NOT NULL,
  `Student_UniversityID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


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



CREATE TABLE `StudentClass` (
  `s_id` int(4) NOT NULL,
  `c_id` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `Teacher` (
  `Id` int(11) NOT NULL,
  `Name` varchar(40) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



CREATE TABLE `TeacherClass` (
  `t_id` int(4) NOT NULL,
  `c_id` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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

ALTER TABLE `Class`
  ADD PRIMARY KEY (`Id`);


ALTER TABLE `Department`
  ADD PRIMARY KEY (`Id`);


ALTER TABLE `Student`
  ADD PRIMARY KEY (`Student_ID`),
  ADD KEY `Student_UniversityID` (`Student_UniversityID`);


ALTER TABLE `StudentClass`
  ADD KEY `student-class_ibfk_1` (`s_id`),
  ADD KEY `student-class_ibfk_2` (`c_id`);


ALTER TABLE `Teacher`
  ADD PRIMARY KEY (`Id`);


ALTER TABLE `TeacherClass`
  ADD KEY `teacher-class_ibfk_1` (`t_id`),
  ADD KEY `teacher-class_ibfk_2` (`c_id`);


ALTER TABLE `University`
  ADD PRIMARY KEY (`University_ID`);


ALTER TABLE `Class`
  MODIFY `Id` int(5) NOT NULL AUTO_INCREMENT;


ALTER TABLE `Department`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;


ALTER TABLE `Student`
  MODIFY `Student_ID` int(11) NOT NULL AUTO_INCREMENT;


ALTER TABLE `Teacher`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;


ALTER TABLE `University`
  MODIFY `University_ID` int(11) NOT NULL AUTO_INCREMENT;


ALTER TABLE `Student`
  ADD CONSTRAINT `student_ibfk_1` FOREIGN KEY (`Student_UniversityID`) REFERENCES `University` (`University_ID`) ON DELETE CASCADE ON UPDATE CASCADE;


ALTER TABLE `StudentClass`
  ADD CONSTRAINT `student-class_ibfk_1` FOREIGN KEY (`s_id`) REFERENCES `Student` (`Student_ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `student-class_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `Class` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;


ALTER TABLE `TeacherClass`
  ADD CONSTRAINT `teacher-class_ibfk_1` FOREIGN KEY (`t_id`) REFERENCES `Teacher` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `teacher-class_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `Class` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE;

