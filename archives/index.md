---
title: 时序
layout: page
---

<ul class="posts">
	  {% for post in site.posts %}
		{% unless post.next %}
		  <h3>{{ post.date | date: '%Y' }}</h3>
		{% else %}
		  {% capture year %}{{ post.date | date: '%Y' }}{% endcapture %}
		  {% capture nyear %}{{ post.next.date | date: '%Y' }}{% endcapture %}
		  {% if year != nyear %}
			<h3>{{ post.date | date: '%Y' }}</h3>
		  {% endif %}
		{% endunless %}
		
		<li><span>{{ post.date | date: "%Y-%m-%d"}}</span> &raquo; <a href="{{ post.url }}">{{ post.title }}</a></li>
	  {% endfor %}
</ul>