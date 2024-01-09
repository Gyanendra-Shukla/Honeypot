function showContent(contentId) {
    const mainContent = document.getElementById('main-content');
    const homeContent = document.getElementById('home-content');
    const aboutContent = document.getElementById('about-content');
    const servicesContent = document.getElementById('services-content');
    const contactContent = document.getElementById('contact-content');

    mainContent.style.display = 'none';
    homeContent.style.display = 'none';
    aboutContent.style.display = 'none';
    servicesContent.style.display = 'none';
    contactContent.style.display = 'none';

    if (contentId === 'home') {
        homeContent.style.display = 'block';
    } else if (contentId === 'about') {
        aboutContent.style.display = 'block';
    } else if (contentId === 'services') {
        servicesContent.style.display = 'block';
    } else if (contentId === 'contact') {
        contactContent.style.display = 'block';
    }
}

