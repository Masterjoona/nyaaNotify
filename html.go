package main

import "regexp"

func MatchPosts() ([][]string, []NyaaPost) {
	postRe := regexp.MustCompile(`(?m)<tr class=".{1,10}">
				<td>
					<a href=".{1,10}" title="(.*?)">
						<img src="(\/static\/img\/icons\/nyaa\/.{1,10})" alt=".*?" class="category-icon">
					<\/a>
				<\/td>
				<td colspan="2">
					(?:(?:<a href="(\/view\/\d+)" title="(.*?)">.*?<\/a>)|(?:<a href="\/view\/\d+#comments" class="comments" title="\d+ comment(?:|s)">
						<i class="fa fa-comments-o"><\/i>(\d+)<\/a>
					<a href="(\/view\/\d+)" title="(.*?)">.*?<\/a>))
				<\/td>
				<td class="text-center">
					<a href="(\/download\/\d+\.torrent)"><i class="fa fa-fw fa-download"><\/i><\/a>
					<a href="(magnet:.*?)"><i class="fa fa-fw fa-magnet"><\/i>\<\/a>
				<\/td>
				<td class="text-center">(.*?)<\/td>
				<td class="text-center" data-timestamp="(\d+)">.*?<\/td>

				<td class="text-center">(\d+)<\/td>
				<td class="text-center">(\d+)<\/td>
				<td class="text-center">(\d+)<\/td>
			<\/tr>`)
	content := FetchNyaa(Url)
	matches := postRe.FindAllStringSubmatch(string(content), -1)
	nyaaPosts := make([]NyaaPost, len(matches))
	return matches, nyaaPosts
}
