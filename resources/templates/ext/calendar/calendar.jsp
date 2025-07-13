<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<html>
<head>
    <title></title>
    <link rel="stylesheet" href="/bower_components/bootstrap-calendar/css/calendar.min.css">
    <!-- script type="text/javascript" src="/bower_components/jquery/dist/jquery.min.js"></script -->
    <script type="text/javascript" src="/bower_components/underscore/underscore-min.js"></script>
    <script type="text/javascript" src="/bower_components/bootstrap-calendar/js/calendar.js"></script>
    <script type="text/javascript" src="/bower_components/bootstrap-calendar/js/language/ko-KR.js"></script>
    <script type="text/javascript">
        $(document).ready(function(){
            $('.btn-group button[data-calendar-nav]').each(function() {
                var $this = $(this);
                $this.click(function() {
                    calendar.navigate($this.data('calendar-nav'));
                });
            });

            $('.btn-group button[data-calendar-view]').each(function() {
                var $this = $(this);
                $this.click(function() {
                    calendar.view($this.data('calendar-view'));
                });
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

    <div class="page-header">

        <div class="pull-right form-inline">
            <div class="btn-group">
                <button class="btn btn-primary" data-calendar-nav="prev"><< 이전</button>
                <button class="btn" data-calendar-nav="today">오늘</button>
                <button class="btn btn-primary" data-calendar-nav="next">다음 >></button>
            </div>
            <div class="btn-group">
                <button class="btn btn-warning" data-calendar-view="year">연간</button>
                <button class="btn btn-warning active" data-calendar-view="month">월별</button>
                <button class="btn btn-warning" data-calendar-view="week">주별</button>
                <button class="btn btn-warning" data-calendar-view="day">일별</button>
            </div>
        </div>

        <h3></h3>
    </div>

    <div id="calendar"></div>
    <script type="text/javascript">
        var calendar = $("#calendar").calendar({
            language: 'ko-KR',
            tmpl_path: "/bower_components/bootstrap-calendar/tmpls/",
            events_source: "/bbs/ext/calendar/contents",
            onAfterViewLoad: function(view) {
                $('.page-header h3').text(this.getTitle());
                $('.btn-group button').removeClass('active');
                $('button[data-calendar-view="' + view + '"]').addClass('active');
            }
        });
    </script>

</body>
</html>