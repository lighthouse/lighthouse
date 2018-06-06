$(document).ready(function(){

/** ===========================================
    Hide / show the master navigation menu
============================================ */

  // console.log('Window Height is: ' + $(window).height());
  // console.log('Document Height is: ' + $(document).height());

  var previousScroll = 0;

  $(window).scroll(function(){

    var currentScroll = $(this).scrollTop();

    /*
      If the current scroll position is greater than 0 (the top) AND the current scroll position is less than the document height minus the window height (the bottom) run the navigation if/else statement.
    */
    if (currentScroll > 0 && currentScroll < $(document).height() - $(window).height()){
      /*
        If the current scroll is greater than the previous scroll (i.e we're scrolling down the page), hide the nav.
      */
      if (currentScroll > previousScroll){
        window.setTimeout(hideNav, 250);
      /*
        Else we are scrolling up (i.e the previous scroll is greater than the current scroll), so show the nav.
      */
      } else {
        window.setTimeout(showNav, 250);
      }
      /* 
        Set the previous scroll value equal to the current scroll.
      */
      previousScroll = currentScroll;
    }

  });

  function hideNav() {
    $("[data-nav-status='toggle']").removeClass("is-hidden").addClass("is-visible");
  }
  function showNav() {
    $("[data-nav-status='toggle']").removeClass("is-visible").addClass("is-hidden");
  }

});