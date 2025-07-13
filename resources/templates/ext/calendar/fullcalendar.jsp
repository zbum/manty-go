<!DOCTYPE html>
<%@ page contentType="text/html; charset=utf-8"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>
<%@ taglib prefix="fmt" uri="http://java.sun.com/jsp/jstl/fmt"%>
<%@ taglib prefix="fn" uri="http://java.sun.com/jsp/jstl/functions" %>
<%@ taglib prefix="ui" uri="http://aquascuba.com/ctl/ui"%>
<%@ taglib prefix="form" uri="http://www.springframework.org/tags/form" %>
<%@ taglib prefix="spring" uri="http://www.springframework.org/tags"%>
<%@ taglib prefix="sec" uri="http://www.springframework.org/security/tags" %>
<%@ page isELIgnored="false"%>
<html>
<head>
    <link rel='stylesheet' href='https://cdnjs.cloudflare.com/ajax/libs/fullcalendar/3.5.0/fullcalendar.css' />
    <script src='/bower_components/jquery/dist/jquery.min.js'></script>
    <script src='/bower_components/moment/min/moment.min.js'></script>
    <script src='/bower_components/fullcalendar/dist/fullcalendar.js'></script>
    <script src='/bower_components/fullcalendar/dist/locale/ko.js'></script>


    <script>
        $(document).ready(function() {
            $('#calendar').fullCalendar({
                locale: 'ko',
                themeSystem: 'bootstrap3',
                events: '/bbs/ext/fullcalendar/contents'
            });

            $("#calendar-button").click(function(){
                location.href ="/bbs/list/calendar";
            });

        });
    </script>
</head>
<body>
    <div class="row">
        <div class="col-md-8">
            <h3>
                <i class="fa fa-calendar"></i>
                만티스쿠바 일정  <span class="hidden-sm hidden-xs">|</span> <small class="hidden-sm hidden-xs">만티스쿠바 일정표</small>
            </h3>
        </div>
    </div>
    <br/>
    <div class="row">
        <div class="col-md-12">
            <div id='calendar'></div>
        </div>
    </div>
    <br/>
    <sec:authorize access="hasRole('ROLE_ADMIN')">
        <div class="row">
            <div class="col-md-7"></div>
            <div class="col-md-5 text-right">
                <button type="button" id="calendar-button" class="btn btn-default">일정관리</button>
            </div>
        </div>
    </sec:authorize>
</body>
</html>
